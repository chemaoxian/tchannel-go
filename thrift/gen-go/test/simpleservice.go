// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package test

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type SimpleService interface {
	// Parameters:
	//  - Arg
	Call(arg *Data) (r *Data, err error)
	Simple() (err error)
}

type SimpleServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewSimpleServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SimpleServiceClient {
	return &SimpleServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewSimpleServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SimpleServiceClient {
	return &SimpleServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Arg
func (p *SimpleServiceClient) Call(arg *Data) (r *Data, err error) {
	if err = p.sendCall(arg); err != nil {
		return
	}
	return p.recvCall()
}

func (p *SimpleServiceClient) sendCall(arg *Data) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Call", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SimpleServiceCallArgs{
		Arg: arg,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *SimpleServiceClient) recvCall() (value *Data, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "Call" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Call failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Call failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Call failed: invalid message type")
		return
	}
	result := SimpleServiceCallResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

func (p *SimpleServiceClient) Simple() (err error) {
	if err = p.sendSimple(); err != nil {
		return
	}
	return p.recvSimple()
}

func (p *SimpleServiceClient) sendSimple() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Simple", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SimpleServiceSimpleArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *SimpleServiceClient) recvSimple() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "Simple" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Simple failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Simple failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Simple failed: invalid message type")
		return
	}
	result := SimpleServiceSimpleResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.SimpleErr != nil {
		err = result.SimpleErr
		return
	}
	return
}

type SimpleServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      SimpleService
}

func (p *SimpleServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SimpleServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SimpleServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSimpleServiceProcessor(handler SimpleService) *SimpleServiceProcessor {

	self4 := &SimpleServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["Call"] = &simpleServiceProcessorCall{handler: handler}
	self4.processorMap["Simple"] = &simpleServiceProcessorSimple{handler: handler}
	return self4
}

func (p *SimpleServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type simpleServiceProcessorCall struct {
	handler SimpleService
}

func (p *simpleServiceProcessorCall) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SimpleServiceCallArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Call", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SimpleServiceCallResult{}
	var retval *Data
	var err2 error
	if retval, err2 = p.handler.Call(args.Arg); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Call: "+err2.Error())
		oprot.WriteMessageBegin("Call", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Call", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type simpleServiceProcessorSimple struct {
	handler SimpleService
}

func (p *simpleServiceProcessorSimple) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SimpleServiceSimpleArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Simple", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SimpleServiceSimpleResult{}
	var err2 error
	if err2 = p.handler.Simple(); err2 != nil {
		switch v := err2.(type) {
		case *SimpleErr:
			result.SimpleErr = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Simple: "+err2.Error())
			oprot.WriteMessageBegin("Simple", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("Simple", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Arg
type SimpleServiceCallArgs struct {
	Arg *Data `thrift:"arg,1" json:"arg"`
}

func NewSimpleServiceCallArgs() *SimpleServiceCallArgs {
	return &SimpleServiceCallArgs{}
}

var SimpleServiceCallArgs_Arg_DEFAULT *Data

func (p *SimpleServiceCallArgs) GetArg() *Data {
	if !p.IsSetArg() {
		return SimpleServiceCallArgs_Arg_DEFAULT
	}
	return p.Arg
}
func (p *SimpleServiceCallArgs) IsSetArg() bool {
	return p.Arg != nil
}

func (p *SimpleServiceCallArgs) Read(iprot thrift.TProtocol) error {
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
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
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
	return nil
}

func (p *SimpleServiceCallArgs) readField1(iprot thrift.TProtocol) error {
	p.Arg = &Data{}
	if err := p.Arg.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Arg), err)
	}
	return nil
}

func (p *SimpleServiceCallArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Call_args"); err != nil {
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

func (p *SimpleServiceCallArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("arg", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:arg: ", p), err)
	}
	if err := p.Arg.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Arg), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:arg: ", p), err)
	}
	return err
}

func (p *SimpleServiceCallArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SimpleServiceCallArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SimpleServiceCallResult struct {
	Success *Data `thrift:"success,0" json:"success,omitempty"`
}

func NewSimpleServiceCallResult() *SimpleServiceCallResult {
	return &SimpleServiceCallResult{}
}

var SimpleServiceCallResult_Success_DEFAULT *Data

func (p *SimpleServiceCallResult) GetSuccess() *Data {
	if !p.IsSetSuccess() {
		return SimpleServiceCallResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SimpleServiceCallResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SimpleServiceCallResult) Read(iprot thrift.TProtocol) error {
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
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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
	return nil
}

func (p *SimpleServiceCallResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &Data{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *SimpleServiceCallResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Call_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *SimpleServiceCallResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *SimpleServiceCallResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SimpleServiceCallResult(%+v)", *p)
}

type SimpleServiceSimpleArgs struct {
}

func NewSimpleServiceSimpleArgs() *SimpleServiceSimpleArgs {
	return &SimpleServiceSimpleArgs{}
}

func (p *SimpleServiceSimpleArgs) Read(iprot thrift.TProtocol) error {
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

func (p *SimpleServiceSimpleArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Simple_args"); err != nil {
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

func (p *SimpleServiceSimpleArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SimpleServiceSimpleArgs(%+v)", *p)
}

// Attributes:
//  - SimpleErr
type SimpleServiceSimpleResult struct {
	SimpleErr *SimpleErr `thrift:"simpleErr,1" json:"simpleErr,omitempty"`
}

func NewSimpleServiceSimpleResult() *SimpleServiceSimpleResult {
	return &SimpleServiceSimpleResult{}
}

var SimpleServiceSimpleResult_SimpleErr_DEFAULT *SimpleErr

func (p *SimpleServiceSimpleResult) GetSimpleErr() *SimpleErr {
	if !p.IsSetSimpleErr() {
		return SimpleServiceSimpleResult_SimpleErr_DEFAULT
	}
	return p.SimpleErr
}
func (p *SimpleServiceSimpleResult) IsSetSimpleErr() bool {
	return p.SimpleErr != nil
}

func (p *SimpleServiceSimpleResult) Read(iprot thrift.TProtocol) error {
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
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
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
	return nil
}

func (p *SimpleServiceSimpleResult) readField1(iprot thrift.TProtocol) error {
	p.SimpleErr = &SimpleErr{}
	if err := p.SimpleErr.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.SimpleErr), err)
	}
	return nil
}

func (p *SimpleServiceSimpleResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Simple_result"); err != nil {
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

func (p *SimpleServiceSimpleResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetSimpleErr() {
		if err := oprot.WriteFieldBegin("simpleErr", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:simpleErr: ", p), err)
		}
		if err := p.SimpleErr.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.SimpleErr), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:simpleErr: ", p), err)
		}
	}
	return err
}

func (p *SimpleServiceSimpleResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SimpleServiceSimpleResult(%+v)", *p)
}
