//go:build gomock || generate

package ackhandler

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\"  -package ackhandler -destination mock_sent_packet_tracker_test.go github.com/quic-go/quic-go/public/ackhandler SentPacketTracker"
type SentPacketTracker = sentPacketTracker

//go:generate sh -c "go run go.uber.org/mock/mockgen -build_flags=\"-tags=gomock\"  -package ackhandler -destination mock_ecn_handler_test.go github.com/quic-go/quic-go/public/ackhandler ECNHandler"
type ECNHandler = ecnHandler
