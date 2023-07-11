// Code generated by thriftgo (0.2.11). DO NOT EDIT.

package idlmanagement

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type IDLManagement interface {
	GetThriftFile(ctx context.Context, fileName string) (r string, err error)
}

type IDLManagementClient struct {
	c thrift.TClient
}

func NewIDLManagementClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *IDLManagementClient {
	return &IDLManagementClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewIDLManagementClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *IDLManagementClient {
	return &IDLManagementClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewIDLManagementClient(c thrift.TClient) *IDLManagementClient {
	return &IDLManagementClient{
		c: c,
	}
}

func (p *IDLManagementClient) Client_() thrift.TClient {
	return p.c
}

func (p *IDLManagementClient) GetThriftFile(ctx context.Context, fileName string) (r string, err error) {
	var _args IDLManagementGetThriftFileArgs
	_args.FileName = fileName
	var _result IDLManagementGetThriftFileResult
	if err = p.Client_().Call(ctx, "GetThriftFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type IDLManagementProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      IDLManagement
}

func (p *IDLManagementProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *IDLManagementProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *IDLManagementProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewIDLManagementProcessor(handler IDLManagement) *IDLManagementProcessor {
	self := &IDLManagementProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("GetThriftFile", &iDLManagementProcessorGetThriftFile{handler: handler})
	return self
}
func (p *IDLManagementProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type iDLManagementProcessorGetThriftFile struct {
	handler IDLManagement
}

func (p *iDLManagementProcessorGetThriftFile) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := IDLManagementGetThriftFileArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetThriftFile", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := IDLManagementGetThriftFileResult{}
	var retval string
	if retval, err2 = p.handler.GetThriftFile(ctx, args.FileName); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetThriftFile: "+err2.Error())
		oprot.WriteMessageBegin("GetThriftFile", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("GetThriftFile", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type IDLManagementGetThriftFileArgs struct {
	FileName string `thrift:"fileName,1" frugal:"1,default,string" json:"fileName"`
}

func NewIDLManagementGetThriftFileArgs() *IDLManagementGetThriftFileArgs {
	return &IDLManagementGetThriftFileArgs{}
}

func (p *IDLManagementGetThriftFileArgs) InitDefault() {
	*p = IDLManagementGetThriftFileArgs{}
}

func (p *IDLManagementGetThriftFileArgs) GetFileName() (v string) {
	return p.FileName
}
func (p *IDLManagementGetThriftFileArgs) SetFileName(val string) {
	p.FileName = val
}

var fieldIDToName_IDLManagementGetThriftFileArgs = map[int16]string{
	1: "fileName",
}

func (p *IDLManagementGetThriftFileArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_IDLManagementGetThriftFileArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *IDLManagementGetThriftFileArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.FileName = v
	}
	return nil
}

func (p *IDLManagementGetThriftFileArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetThriftFile_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *IDLManagementGetThriftFileArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("fileName", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.FileName); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *IDLManagementGetThriftFileArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IDLManagementGetThriftFileArgs(%+v)", *p)
}

func (p *IDLManagementGetThriftFileArgs) DeepEqual(ano *IDLManagementGetThriftFileArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.FileName) {
		return false
	}
	return true
}

func (p *IDLManagementGetThriftFileArgs) Field1DeepEqual(src string) bool {

	if strings.Compare(p.FileName, src) != 0 {
		return false
	}
	return true
}

type IDLManagementGetThriftFileResult struct {
	Success *string `thrift:"success,0,optional" frugal:"0,optional,string" json:"success,omitempty"`
}

func NewIDLManagementGetThriftFileResult() *IDLManagementGetThriftFileResult {
	return &IDLManagementGetThriftFileResult{}
}

func (p *IDLManagementGetThriftFileResult) InitDefault() {
	*p = IDLManagementGetThriftFileResult{}
}

var IDLManagementGetThriftFileResult_Success_DEFAULT string

func (p *IDLManagementGetThriftFileResult) GetSuccess() (v string) {
	if !p.IsSetSuccess() {
		return IDLManagementGetThriftFileResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *IDLManagementGetThriftFileResult) SetSuccess(x interface{}) {
	p.Success = x.(*string)
}

var fieldIDToName_IDLManagementGetThriftFileResult = map[int16]string{
	0: "success",
}

func (p *IDLManagementGetThriftFileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IDLManagementGetThriftFileResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_IDLManagementGetThriftFileResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *IDLManagementGetThriftFileResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Success = &v
	}
	return nil
}

func (p *IDLManagementGetThriftFileResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetThriftFile_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *IDLManagementGetThriftFileResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Success); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *IDLManagementGetThriftFileResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IDLManagementGetThriftFileResult(%+v)", *p)
}

func (p *IDLManagementGetThriftFileResult) DeepEqual(ano *IDLManagementGetThriftFileResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *IDLManagementGetThriftFileResult) Field0DeepEqual(src *string) bool {

	if p.Success == src {
		return true
	} else if p.Success == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Success, *src) != 0 {
		return false
	}
	return true
}
