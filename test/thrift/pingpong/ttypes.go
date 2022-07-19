// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package pingpong

import (
	"bytes"
	"fmt"

	"github.com/uber/tchannel-go/thirdparty/github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - Key
type Ping struct {
	Key string `thrift:"key,1,required" json:"key"`
}

func NewPing() *Ping {
	return &Ping{}
}

func (p *Ping) GetKey() string {
	return p.Key
}
func (p *Ping) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetKey bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetKey = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetKey {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Key is not set"))
	}
	return nil
}

func (p *Ping) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Key = v
	}
	return nil
}

func (p *Ping) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Ping) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
	}
	if err := oprot.WriteString(string(p.Key)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
	}
	return err
}

func (p *Ping) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Ping(%+v)", *p)
}

// Attributes:
//  - Source
type Pong struct {
	Source string `thrift:"source,1,required" json:"source"`
}

func NewPong() *Pong {
	return &Pong{}
}

func (p *Pong) GetSource() string {
	return p.Source
}
func (p *Pong) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetSource bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetSource = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetSource {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Source is not set"))
	}
	return nil
}

func (p *Pong) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Source = v
	}
	return nil
}

func (p *Pong) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("pong"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Pong) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("source", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:source: ", p), err)
	}
	if err := oprot.WriteString(string(p.Source)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.source (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:source: ", p), err)
	}
	return err
}

func (p *Pong) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Pong(%+v)", *p)
}

type PingError struct {
}

func NewPingError() *PingError {
	return &PingError{}
}

func (p *PingError) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PingError) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("PingError"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PingError) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PingError(%+v)", *p)
}

func (p *PingError) Error() string {
	return p.String()
}
