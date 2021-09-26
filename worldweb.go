package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	types "github.com/gfx-labs/etherlands/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type DistrictMetadata struct {
	Owner       string            `json:"owner"`
	Name        string            `json:"name"`
	Contains    []uint64          `json:"contains"`
	Clusters    []ClusterMetadata `json:"clusters"`
	Description string            `json:"description"`
	Image       string            `json:"image"`
	ExternalURL string            `json:"external_url"`
	Attributes  []Attribute       `json:"attributes"`
}

type Attribute struct {
	DisplayType string `json:"display_type,omitempty"`
	TraitType   string `json:"trait_type"`
	Value       int64  `json:"value"`
}

type PlotSearchResult struct {
	IdArray       []uint64   `json:"id_array"`
	DistrictArray []uint64   `json:"district_array"`
	LocationArray [][2]int64 `json:"location_array"`

	Count int `json:"count"`
}

func sendFail(w http.ResponseWriter, err error) bool {
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return true
	}
	return false
}

const LIMIT = 1000 * 1000

func (WW *worldweb) serveNftImage(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	id := chi.URLParam(r, "id")
	location := path.Join("./db", "images", "opensea", slug, id)
	if _, err := os.Stat(location); err == nil {
		image, err := ioutil.ReadFile(location)
		if err == nil {
			w.WriteHeader(200)
			w.Write(image)
			return
		}
	}
	w.WriteHeader(400)
	w.Write([]byte("image not cached"))
}

func (WW *worldweb) serve24Creator(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	encodedname := types.Create24Name(name)
	hex_name := hexutil.Encode(encodedname[:])
	w.WriteHeader(200)
	w.Write([]byte(hex_name))
}

func (WW *worldweb) servePlotQuery(w http.ResponseWriter, r *http.Request) {
	x1_string := chi.URLParam(r, "x1")
	x2_string := chi.URLParam(r, "x2")
	z1_string := chi.URLParam(r, "z1")
	z2_string := chi.URLParam(r, "z2")
	x1, err := strconv.ParseInt(x1_string, 10, 64)
	if sendFail(w, err) {
		return
	}
	x2, err := strconv.ParseInt(x2_string, 10, 64)
	if sendFail(w, err) {
		return
	}
	z1, err := strconv.ParseInt(z1_string, 10, 64)
	if sendFail(w, err) {
		return
	}
	z2, err := strconv.ParseInt(z2_string, 10, 64)
	if sendFail(w, err) {
		return
	}
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if z1 > z2 {
		z1, z2 = z2, z1
	}
	if (x2-x1+1)*(z2-z1+1) > LIMIT {
		if sendFail(w, errors.New("query too large")) {
			return
		}
	}
	id_array := []uint64{}
	district_array := []uint64{}
	location_array := [][2]int64{}
	for x := x1; x <= x2; x++ {
		for z := z1; z <= z2; z++ {
			plot_id, err := WW.W.SearchPlot(x, z)
			if err == nil {
				id_array = append(id_array, plot_id.PlotId())
				district_array = append(district_array, plot_id.DistrictId())
				location_array = append(location_array, [2]int64{plot_id.X(), plot_id.Z()})
			}
		}
	}
	pending, err := json.Marshal(PlotSearchResult{
		IdArray:       id_array,
		DistrictArray: district_array,
		LocationArray: location_array,
		Count:         len(id_array),
	})
	if sendFail(w, err) {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(pending)
}

func (WW *worldweb) serveDistrictMetadata(
	w http.ResponseWriter,
	r *http.Request,
) {
	id_string := chi.URLParam(r, "id")
	district_id, err := strconv.ParseUint(id_string, 10, 64)
	if sendFail(w, err) {
		return
	}

	district, err := WW.W.GetDistrict(district_id)
	if sendFail(w, err) {
		return
	}

	count := WW.W.PlotsOfDistrict(district_id)
	clustered := GenerateClusterMetadata(WW.W, count)

	output := fmt.Sprintf("A District containing %d plots at locations", len(count))
	for _, v := range count {
		plot, err := WW.W.GetPlot(v)
		if err == nil {
			output = fmt.Sprintf(output+", [%d,%d]", plot.X(), plot.Z())
		}
	}
	district_attr := []Attribute{}
	district_attr = append(district_attr,
		Attribute{
			DisplayType: "number",
			TraitType:   "Size",
			Value:       int64(len(count)),
		})
	metadata := DistrictMetadata{
		Owner:       district.OwnerAddress(),
		Name:        district.StringName(),
		Contains:    count,
		Clusters:    clustered,
		Description: output,
		Image:       "https://i.imgur.com/TZKmzvw.png",
		ExternalURL: fmt.Sprintf("https://etherlands.com/district/%d", district.DistrictId()),
		Attributes:  district_attr,
	}

	//DrawDistrict(metadata)
	pending, err := json.Marshal(metadata)
	if sendFail(w, err) {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(pending)
}

func (WW *worldweb) serveLinkForwarder(
	w http.ResponseWriter,
	r *http.Request,
) {
	msg := strings.Replace(chi.URLParam(r, "message"), ":", "", -1)
	sig := strings.Replace(chi.URLParam(r, "signature"), ":", "", -1)
	pkey := strings.Replace(chi.URLParam(r, "publickey"), ":", "", -1)

	tosend := fmt.Sprintf("%s:%s:%s", msg, sig, pkey)

	if (msg != "") && (sig != "") && (pkey != "") {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}
	w.Write([]byte(tosend))
}

type worldweb struct {
	W *types.World //WWW get it....
}

func StartWorldWeb(W *types.World) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	www := &worldweb{
		W: W,
	}
	r.Get("/district/{id}", www.serveDistrictMetadata)
	r.Get("/plot_query/{x1}/{x2}/{z1}/{z2}", www.servePlotQuery)
	r.Get("/encode_ledders/{name}", www.serve24Creator)
	r.Get("/link/{message}/{signature}/{publickey}/{a}/{b}/{c}", www.serveLinkForwarder)
	r.Get("/nft_image/{slug}/{id}", www.serveNftImage)

	log.Println("launching web service at 10100")
	go func() {
		err := http.ListenAndServe(":10100", r)
		if err != nil {
			log.Println("failed to start web service", err)
		}
	}()

}
