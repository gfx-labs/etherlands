package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"log"
	"math"
	"os"
	"sort"
)


const WIDTH = 1000;
const LENGTH = 1000;

const GAP = 1;


func upsideDown(pixels [][]color.Color)  {
  for i:=0; i<len(pixels);i++{
     tr:=pixels[i]
     for j:=0; j<len(tr)/2;j++{
        k:= len(tr) -j -1
        tr[j],tr[k] = tr[k],tr[j]
     }
  }
}


func DrawDistrict(input DistrictMetadata) (error) {
	grid_max_x := int64(0)
	grid_max_z := int64(0)
	shapes := []*ClusterShape{}
	for _, cluster := range input.Clusters {
		grid_max_z = grid_max_z + 1 + cluster.LengthZ;
		grid_max_x = grid_max_x + 1 + cluster.LengthX;
		shapes = append(shapes, CreateClusterShape(cluster))
	}
	display := CreateDistrictDisplay(shapes,grid_max_x , grid_max_z)
	size := 30
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(grid_max_x)*size, int(grid_max_z)*size}

	img:= image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for x := 0; x < int(grid_max_x)*size; x = x + size {
    for z := 0; z < int(grid_max_z)*size; z = z + size {
			check := display.existsAt(int64(x/size),int64(z/size))
			if check != -1 {
				for i:=0; i < size; i++{
					for j:=0; j < size; j++{
						img.Set(x+i,(int(grid_max_z)*size - (z+j) ),palette.Plan9[(check * 20 + 30)])
					}
				}
			}
    }
	}
	f, _ := os.Create("db/test.png")
	png.Encode(f,img)

	total := "\n"
	for i := int64(0); i < display.z_max;i++ {
		line := ""
		for j := int64(0); j < display.x_max; j++ {
			char := "-"
			if display.existsAt(j,i) != -1{
				char = fmt.Sprintf("%d",display.existsAt(j,i))
			}
			line = line + char
		}
		line = line + "\n"
		total = total + line
	}
	log.Println(display)
	log.Println(total)
	return nil;
}

type DistrictDisplay struct {
	x_max int64
	z_max int64

	total_shapes int
	shapes []*ClusterShape
	x_loc []int64
	z_loc []int64
}


func CreateDistrictDisplay(shapes []*ClusterShape, xmax, zmax int64) *DistrictDisplay{
	sort.Slice(shapes, func (i, j int) bool {
		return shapes[i].origin[1]< shapes[j].origin[1]
	})
	D := DistrictDisplay{
		shapes: shapes,
		total_shapes: len(shapes),
		x_max: xmax,
		z_max: zmax,
	}

	for i:= 1; i < D.total_shapes; i++ {
		D.moveDownUntilCan(i)
	}
	sort.Slice(D.shapes, func (i, j int) bool {
		return D.shapes[i].origin[0]< D.shapes[j].origin[0]
	})
	for i:= 1; i < D.total_shapes; i++ {
		D.moveRightUntilCan(i)
	}
	return &D
}


func (D *DistrictDisplay) moveRightUntilCannot(id int) {
	iters := 0;
	done := false
	for !done {
		if iters > int(D.x_max) {
			done = true;
		}
		if D.canMoveClusterHorizontal(id, 1) {
			D.shapes[id].x_loc = D.shapes[id].x_loc + 1
		}else{
			done = true;
		}
		iters = iters + 1
	}
}
func (D *DistrictDisplay) moveRightUntilCan(id int) {
	iters := int64(1);
	done := false
	for !done {
		if iters > D.x_max {
			D.shapes[id].x_loc = D.shapes[id].x_loc + iters + GAP
			done = true;
		}
		if D.canMoveClusterHorizontal(id, iters) {
			D.shapes[id].x_loc = D.shapes[id].x_loc + iters + GAP
			done = true;
		}else{
			iters = iters + 1
		}
	}
}

func (D *DistrictDisplay) moveDownUntilCan(id int) {
	iters := int64(1);
	done := false
	for !done {
		if iters > D.z_max {
			D.shapes[id].z_loc = D.shapes[id].z_loc + iters + GAP
			done = true;
		}
		if D.canMoveClusterVertical(id, iters) {
			D.shapes[id].z_loc = D.shapes[id].z_loc + iters + GAP
			done = true;
		}else{
			iters = iters + 1
		}
	}
}




func (D *DistrictDisplay) existsAt(x, z int64) (int) {

	output:= -1
	for i:=0; i < D.total_shapes; i++{
		x_o := D.shapes[i].x_loc
		z_o := D.shapes[i].z_loc;
		if D.shapes[i].HasAt(x - x_o, z - z_o){
			output= i
		}
	}
	return output
}

func (D *DistrictDisplay) existsAtExcept(id int, x, z int64) (int) {
	output:= -1
	for i:=0; i < D.total_shapes; i++{
		if(id != i){
			x_o := D.shapes[i].x_loc
			z_o := D.shapes[i].z_loc;
			if D.shapes[i].HasAt(x - x_o, z - z_o){
				output= i
			}
		}
	}
	return output
}

func (D *DistrictDisplay) canMoveClusterHorizontal(id int, pixels int64) bool {
	target := D.shapes[id]

	if (target.x_loc + pixels) > D.x_max {
		return false;
	}
	lower_extremities := make([]int64, target.z_max)
	upper_extremities := make([]int64, target.z_max)
	for i := range lower_extremities{
		lower_extremities[i] = math.MaxInt64
	}
	for i := int64(0); i < target.count; i++{
		if lower_extremities[target.coords[i][1]] > target.coords[i][0] {
			lower_extremities[target.coords[i][1]] = target.coords[i][0]
		}
		if upper_extremities[target.coords[i][1]] < target.coords[i][0] {
			upper_extremities[target.coords[i][1]] = target.coords[i][0]
		}
	}
	for j:=0; j < int(target.z_max); j++{
		for k := lower_extremities[j]; k <= (upper_extremities[j]+GAP); k++ {
			check := D.existsAtExcept(id, target.x_loc + k + pixels, target.z_loc)
			if(check != -1){
				return false
			}
		}
	}

	return true
}

func (D *DistrictDisplay) canMoveClusterVertical(id int, pixels int64) bool {
	target := D.shapes[id]

	if (target.z_loc + pixels) > D.z_max {
		return false;
	}
	lower_extremities := make([]int64, target.x_max)
	upper_extremities := make([]int64, target.x_max)
	for i := range lower_extremities{
		lower_extremities[i] = math.MaxInt64
	}
	for i := int64(0); i < target.count; i++{
		if lower_extremities[target.coords[i][0]] > target.coords[i][1] {
			lower_extremities[target.coords[i][0]] = target.coords[i][1]
		}
		if upper_extremities[target.coords[i][0]] < target.coords[i][1] {
			upper_extremities[target.coords[i][0]] = target.coords[i][1]
		}
	}
	for j:=0; j < int(target.x_max); j++{
		for k := lower_extremities[j]; k <= upper_extremities[j] + GAP; k++ {
			check := D.existsAtExcept(id, target.x_loc, k + pixels + target.z_loc)
			if(check != -1){
				return false
			}
		}
	}

	return true
}

type ClusterShape struct {
	coord_map map[[2]int64]bool
	x_max int64
	z_max int64

	x_loc int64
	z_loc int64



	coords [][2]int64
	origin [2]int64
	count int64
}

func (S *ClusterShape) HasAt(x, z int64) bool {
	v1, ok1 := S.coord_map[[2]int64{x,z}];
	if(ok1){
		return v1;
	}
	return false
}


func CreateClusterShape(cluster ClusterMetadata) *ClusterShape{
	shape := &ClusterShape{
		coord_map: make(map[[2]int64]bool),
		x_loc: int64(0),
		z_loc: int64(0),
		x_max: cluster.LengthX,
		z_max: cluster.LengthZ,
		coords: cluster.Offsets,
		origin: [2]int64{cluster.OriginX,cluster.OriginZ},
		count: int64(len(cluster.PlotIds)),
	}
	for _, coord := range shape.coords {
		shape.coord_map[coord] = true;
	}
	return shape
}

type byCoord []*ClusterShape

func (s byCoord) Len() int {
	return len(s)
}
func (s byCoord) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
