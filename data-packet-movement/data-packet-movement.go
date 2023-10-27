type DecodedPacket struct {
    Short1 int16
    Char12 string
    Byte1 byte
    Char8 string
    Short2 int16
    Char15 string
    Long4 int32
}

func DecodePacket(packet []byte) DecodedPacket {
    var decodedPacket DecodedPacket
    decodedPacket.Short1 = int16(binary.BigEndian.Uint16(packet[0:2]))
    decodedPacket.Char12 = string(packet[2:14])
    decodedPacket.Byte1 = packet[14]
    decodedPacket.Char8 = string(packet[15:23])
    decodedPacket.Short2 = int16(binary.BigEndian.Uint16(packet[23:25]))
    decodedPacket.Char15 = string(packet[25:40])
    decodedPacket.Long4 = int32(binary.BigEndian.Uint32(packet[40:44]))
    return decodedPacket
}
