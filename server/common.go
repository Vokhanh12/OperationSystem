package main

import "math"

type TamGiac struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
}

type TamGiac_kq struct {
	Cv float64 `json:"cv"`
	Dt float64 `json:"dt"`
}

func (tg TamGiac) TinhChuVi() float64 {

	return tg.A + tg.B + tg.C

}

// Diện tích Heron (hoặc Diện tích Heron)
func (tg TamGiac) TinhDienTich() float64 {

	// S= p(p−a)(p−b)(p−c)

	// Kiểm tra xem tam giác có hợp lệ không
	if tg.A <= 0 || tg.B <= 0 || tg.C <= 0 || (tg.A+tg.B <= tg.C) || (tg.A+tg.C <= tg.B) || (tg.B+tg.C <= tg.A) {
		return -1 // Trả về -1 nếu tam giác không hợp lệ
	}

	// Tính toán diện tích bằng công thức Heron
	p := (tg.A + tg.B + tg.C) / 2
	return math.Sqrt(p * (p - tg.A) * (p - tg.B) * (p - tg.C))

}
