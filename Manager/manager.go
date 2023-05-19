package Manager

var idC int64 = 0

func GetNewCouponId() int64 {
	idC++
	return idC
}
