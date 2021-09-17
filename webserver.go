package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)


type DistrictMetadata struct {
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Contains    []uint64   `json:"contains"`
	Clusters    []ClusterMetadata `json:"clusters"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	ExternalURL string  `json:"external_url"`
	Attributes  []Attribute `json:"attributes"`
}

type Attribute  struct {
	DisplayType string `json:"display_type,omitempty"`
	TraitType   string `json:"trait_type"`
	Value       int64    `json:"value"`
}

type PlotSearchResult struct {
	IdArray []uint64 `json:"id_array"`
	DistrictArray []uint64 `json:"district_array"`
	LocationArray [][2]int64 `json:"location_array"`

	Count int `json:"count"`
}


func sendFail(w http.ResponseWriter, err error) bool{
	if(err != nil){
		w.WriteHeader(200)
		w.Write([]byte(err.Error()))
		return true
	}
	return false
}

const LIMIT = 1000*1000


func (E *EtherlandsContext) ServePlotQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	x1_string := ps.ByName("x1")
	x2_string := ps.ByName("x2")
	z1_string := ps.ByName("z1")
	z2_string := ps.ByName("z2")
	x1, err := strconv.ParseInt(x1_string,10,64)
	if sendFail(w, err) {return}
	x2, err := strconv.ParseInt(x2_string,10,64)
	if sendFail(w, err) {return}
	z1, err := strconv.ParseInt(z1_string,10,64)
	if sendFail(w, err) {return}
	z2, err := strconv.ParseInt(z2_string,10,64)
	if sendFail(w, err) {return}
	if x1 > x2 {
		x1, x2 = x2, x1;
	}
	if z1 > z2 {
		z1, z2 = z2, z1;
	}
	if (x2-x1+1)*(z2-z1+1) > LIMIT {
		if sendFail(w, errors.New("query too large")) {return}
	}
	id_array := []uint64{}
	district_array := []uint64{}
	location_array := [][2]int64{}
	for x := x1; x <= x2; x++ {
		for z := z1; z <= z2; z++ {
			plot_id, err := E.SearchPlot(x,z)
			if err == nil{
				id_array = append(id_array, plot_id.PlotId())
				district_array = append(district_array, plot_id.DistrictId())
				location_array = append(location_array, [2]int64{plot_id.X(),plot_id.Z()})
			}
		}
	}
	pending, err:= json.Marshal(PlotSearchResult{
		IdArray: id_array,
		DistrictArray: district_array,
		LocationArray: location_array,
		Count: len(id_array),
	})
	if sendFail(w, err) {return}
	w.Header().Add("Content-Type","application/json");
	w.Write(pending)
}

func (E *EtherlandsContext) ServeDistrictMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id_string := ps.ByName("id")
	district_id, err := strconv.ParseUint(id_string,10,64);
	if sendFail(w, err) {return}

	district, err := E.GetDistrict(district_id);
	if sendFail(w, err) {return}

	count := E.plots_zset.GetKeysByScore(district_id)
	clustered := E.GenerateClusterMetadata(count);

	output := fmt.Sprintf("A District containing %d plots at locations", len(count))
	for _, v := range count{
		plot, err := E.GetPlot(v)
		if(err == nil){
			output = fmt.Sprintf(output +", [%d,%d]", plot.X(), plot.Z())
		}
	}
	district_attr := []Attribute{}
	district_attr = append(district_attr,
	Attribute{
		DisplayType: "number",
		TraitType: "Size",
		Value: int64(len(count)),
	})
	metadata := DistrictMetadata{
		Owner: district.OwnerAddress(),
		Name: district.Nickname(),
		Contains: count,
		Clusters: clustered,
		Description: output,
		Image:"https://i.imgur.com/TZKmzvw.png",
		ExternalURL: fmt.Sprintf("https://etherlands.io/district/%d",district.DistrictId()),
		Attributes: district_attr,
	}
	pending, err:= json.Marshal(metadata)
	if sendFail(w,err) {return;}
	w.Header().Add("Content-Type","application/json");
	w.Write(pending)
}




func (E *EtherlandsContext) StartWebService() {
	router := httprouter.New()
	router.GET("/district/:id", E.ServeDistrictMetadata)
	router.GET("/plot_query/:x1/:x2/:z1/:z2", E.ServePlotQuery)
	log.Println("now hosting web service at 10100")
	http.ListenAndServe(":10100", router)
}
