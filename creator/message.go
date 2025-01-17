package creator

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

// Message format for kafka message
type Message struct {
	Sequence    uint64
	ID          string
	CreatedTime time.Time
	Data        []byte
}

// Bytes Serializes the message via gob encoder
func (m Message) Bytes() ([]byte, error) {
	var messageBuffer bytes.Buffer
	enc := gob.NewEncoder(&messageBuffer)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	return messageBuffer.Bytes(), nil
}

// FromBytes DeSerializes the bytes data to a message via gob decoder
func FromBytes(data []byte) (Message, error) {
	var messageBuffer bytes.Buffer
	messageBuffer.Write(data)
	enc := gob.NewDecoder(&messageBuffer)
	var m Message
	err := enc.Decode(&m)
	return m, err
}

func (m Message) String() string {
	return fmt.Sprintf("ID: %s sequence: %d time: %s", m.ID, m.Sequence, m.CreatedTime.Format(time.RFC3339))
}
