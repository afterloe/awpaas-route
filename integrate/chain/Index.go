package chain

type callback func(...interface{}) string

type ChannelChain struct {
	fn callback
	next *ChannelChain
}

func (this *ChannelChain) SetNextSuccess(success *ChannelChain) *ChannelChain {
	this.next = success
	return this.next
}

func (this *ChannelChain) PassChannel(args ...interface{}) string {
	let := this.fn(args...)
	if "next" == let {
		return this.next.PassChannel(args...)
	}
	return let
}