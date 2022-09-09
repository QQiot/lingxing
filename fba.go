package lingxing

// FBA
type fbaService struct {
	Shipment   fbaShipmentService   // FBA 发货单
	StorageFee fbaStorageFeeService // FBA 仓储费
}
