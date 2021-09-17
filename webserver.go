package main

import (
	"encoding/json"
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
	Clusters    [][]uint64 `json:"clusters"`
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

func (E *EtherlandsContext) ServeDistrictMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id_string := ps.ByName("id")
	district_id, err := strconv.ParseUint(id_string,10,64);
	if err != nil {
		w.WriteHeader(200)
		w.Write([]byte("invalid district id provided"))
		return
	}

	district := E.GetDistrict(district_id);
	if district == nil{
		w.WriteHeader(200)
		w.Write([]byte("could not locate district"))
	}



	count := E.plots_zset.GetKeysByScore(district_id)
	clustered := E.Cluster(count);

	output := fmt.Sprintf("A District containing %d plots at locations", len(count))
	for _, v := range count{
		plot := E.GetPlot(v)
		if(plot != nil){
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
	if err != nil{
		w.WriteHeader(200)
		w.Write([]byte("internal error"))
	}
	w.Header().Add("Content-Type","application/json");
	w.Write(pending)
}




func (E *EtherlandsContext) StartWebService() {
	router := httprouter.New()
	router.GET("/district/:id", E.ServeDistrictMetadata)
	log.Println("now hosting web service at 10100")
	http.ListenAndServe(":10100", router)
}
