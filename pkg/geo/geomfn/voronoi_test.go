// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package geomfn

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/geo"
)

func TestVoronoiPolygons(t *testing.T) {

	type args struct {
		a geo.Geometry
		tol float64
		poly int
	}
	tests := []struct {
		name       string
		args      args
		expected   geo.Geometry
		expectedErr error
	}{
		{
			name: "Computes Voronoi Polygons for a given MultiPoint",
			args: args{
				a: geo.MustParseGeometry("POINT(10 20)"),
				tol: 0,
				poly: 1,
			},
			expected: geo.MustParseGeometry("MULTILINESTRING EMPTY"),
		},
		//{
		//	name: "Computes Voronoi Polygons for a given MultiPoint",
		//	args: args{
		//		a: geo.MustParseGeometry("MULTIPOINT ((280 300), (420 330), (380 230), (320 160))"),
		//		tol: 0,
		//		poly: 0,
		//	},
		//	expected: geo.MustParseGeometry("GEOMETRYCOLLECTION (POLYGON ((110 175.71428571428572, 110 500, 310.35714285714283 500, 353.515625 298.59375, 306.875 231.96428571428572, 110 175.71428571428572)), POLYGON ((590 204, 590 -10, 589.1666666666666 -10, 306.875 231.96428571428572, 353.515625 298.59375, 590 204)), POLYGON ((110 -10, 110 175.71428571428572, 306.875 231.96428571428572, 589.1666666666666 -10, 110 -10)), POLYGON ((310.35714285714283 500, 590 500, 590 204, 353.515625 298.59375, 310.35714285714283 500)))"),
		//},
		{
			name: "Computes Voronoi Polygons for a given MultiPoint",
			args: args{
				a: geo.MustParseGeometry("MULTIPOINT ((170 270), (270 270), (230 310), (180 330), (250 340), (315 318), (330 260), (240 170), (220 220), (270 220))"),
				tol: 0,
				poly: 1,
			},
			expected: geo.MustParseGeometry("MULTILINESTRING ((190 510, 213.9473684210526 342.3684210526316), (213.9473684210526 342.3684210526316, 195.625 296.5625), (195.625 296.5625, 0 329.1666666666667), (195.625 296.5625, 216 266), (216 266, 88.33333333333333 138.3333333333333), (88.33333333333333 138.3333333333333, 0 76.50000000000001), (213.9473684210526 342.3684210526316, 267 307), (267 307, 225 265), (225 265, 216 266), (245 245, 225 265), (267 307, 275.9160583941606 309.5474452554744), (275.9160583941606 309.5474452554744, 303.1666666666667 284), (303.1666666666667 284, 296.6666666666667 245), (296.6666666666667 245, 245 245), (245 245, 245 201), (245 201, 88.33333333333333 138.3333333333333), (245 201, 380 120), (380 120, 500 0), (343.7615384615385 510, 275.9160583941606 309.5474452554744), (296.6666666666667 245, 380 120), (500 334.9051724137931, 303.1666666666667 284))"),
		},
		{
			name: "Computes Voronoi Polygons for a given MultiPoint with tolerance",
			args: args{
				a: geo.MustParseGeometry("MULTIPOINT ((150 210), (210 270), (150 220), (220 210), (215 269))"),
				tol: 10,
				poly: 1,
			},
			expected: geo.MustParseGeometry("MULTILINESTRING ((185 215, 187.9268292682927 235.4878048780488), (187.9268292682927 235.4878048780488, 290 252.5), (185 140, 185 215), (185 215, 80 215), (100.8333333333334 340, 187.9268292682927 235.4878048780488))"),
		},
		{
			name: "Computes Voronoi Polygons for a given MultiPoint",
			args: args{
				a: geo.MustParseGeometry("MULTIPOINT ((280 300), (420 330), (380 230), (320 160))"),
				tol: 0,
				poly: 1,
			},
			expected: geo.MustParseGeometry("MULTILINESTRING ((310.3571428571428 500, 353.515625 298.59375), (353.515625 298.59375, 306.875 231.9642857142857), (306.875 231.9642857142857, 110 175.7142857142857), (589.1666666666666 -10, 306.875 231.9642857142857), (353.515625 298.59375, 590 204))"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := VoronoiPolygons(tc.args.a, tc.args.tol, tc.args.poly)
			if tc.expectedErr != nil && tc.expectedErr.Error() != err.Error() {
				t.Errorf("VoronoiPolygons() error = %v, wantErr %v", err, tc.expectedErr)
				return
			}
			require.Equal(t, true, EqualsExact(actual, tc.expected, 1e-10))
		})
	}
}

func TestVoronoiPolygonsWithEnv(t *testing.T) {

	type args struct {
		a geo.Geometry
		b geo.Geometry
		tol float64
		poly int
	}
	tests := []struct {
		name       string
		args      args
		expected   geo.Geometry
		expectedErr error
	}{
		{
			name: "Computes Voronoi Polygons for a given MultiPoint",
			args: args{
				a: geo.MustParseGeometry("POINT(10 20)"),
				b: geo.MustParseGeometry("POLYGON EMPTY"),
				tol: 0,
				poly: 1,
			},
			expected: geo.MustParseGeometry("MULTILINESTRING EMPTY"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := VoronoiPolygonsWithEnv(tc.args.a, tc.args.b, tc.args.tol, tc.args.poly)
			if tc.expectedErr != nil && tc.expectedErr.Error() != err.Error() {
				t.Errorf("VoronoiPolygons() error = %v, wantErr %v", err, tc.expectedErr)
				return
			}
			require.Equal(t, true, EqualsExact(actual, tc.expected, 1e-10))
		})
	}
}
