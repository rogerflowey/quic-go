package ackhandler

import (
	"github.com/quic-go/quic-go/logging"
	"github.com/quic-go/quic-go/public/protocol"
	"github.com/quic-go/quic-go/public/utils"
)

// NewAckHandler creates a new SentPacketHandler and a new ReceivedPacketHandler.
// clientAddressValidated indicates whether the address was validated beforehand by an address validation token.
// clientAddressValidated has no effect for a client.
func NewAckHandler(
	initialPacketNumber protocol.PacketNumber,
	initialMaxDatagramSize protocol.ByteCount,
	rttStats *utils.RTTStats,
	clientAddressValidated bool,
	enableECN bool,
	pers protocol.Perspective,
	tracer *logging.ConnectionTracer,
	logger utils.Logger,
) (SentPacketHandler, ReceivedPacketHandler) {
	sph := newSentPacketHandler(initialPacketNumber, initialMaxDatagramSize, rttStats, clientAddressValidated, enableECN, pers, tracer, logger)
	return sph, newReceivedPacketHandler(sph, logger)
}
