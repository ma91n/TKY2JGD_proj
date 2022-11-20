package main

import (
	"fmt"
	"github.com/everystreet/go-proj/v8/cproj"
	"github.com/everystreet/go-proj/v8/proj"
	"github.com/golang/geo/s1"
)

func main() {
	pj := NewProjection("EPSG:4301", "EPSG:4326")
	defer pj.Close()

	for i := 0; i < 1000*1000; i++ {
		wgsLng, wgsLat := pj.CRSToCRS(128542740/float64(60*60*256), 32706756/float64(60*60*256))
		fmt.Printf("%f %f\n", wgsLng, wgsLat)
	}
}

type Projection struct {
	ctx        *cproj.PJ_CONTEXT
	src        *cproj.PJ
	dst        *cproj.PJ
	pj         *cproj.PJ
	normalized *cproj.PJ
}

func NewProjection(source, target string) Projection {
	ctx := cproj.Context_create()
	src := cproj.Create(ctx, source)
	dst := cproj.Create(ctx, target)
	pj := cproj.Create_crs_to_crs_from_pj(ctx, src, dst, nil, nil)
	normalized := cproj.Normalize_for_visualization(ctx, pj)
	return Projection{
		ctx:        ctx,
		src:        src,
		dst:        dst,
		pj:         pj,
		normalized: normalized,
	}
}

func (p Projection) CRSToCRS(lng, lat float64) (float64, float64) {
	coord := proj.LP{
		Lng: s1.Angle(lng),
		Lat: s1.Angle(lat),
	}
	proj.TransformForward(p.normalized, &coord)
	return float64(coord.Lng), float64(coord.Lat)
}

func (p Projection) Close() {
	cproj.Context_destroy(p.ctx)
	cproj.Destroy(p.src)
	cproj.Destroy(p.dst)
	cproj.Destroy(p.pj)
	cproj.Destroy(p.normalized)
}
