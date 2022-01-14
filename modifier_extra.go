package milter

type ModifierProperties struct {
	WritePacket func(*Message) error
}

func (m *Modifier) SetProperties(properties *ModifierProperties) {
	m.writePacket = properties.WritePacket
}
