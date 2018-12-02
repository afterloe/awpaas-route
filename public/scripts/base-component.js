"use strict";

class ModalWindow_alert extends React.Component {
    constructor(props) {
        super(props);
    }

    closeModal(event) {
        const key = event.target.getAttribute("dataClose") || "";
        if (-1 == key) {
            ReactDOM.unmountComponentAtNode(document.getElementById("modal"));
        }
    }

    render() {
        const {title = "确认删除？", context = "111111"} = this.props;
        return (
            <div className="modal fade bd-example-modal-sm show" tabIndex="-1" dataClose="-1" onClick={this.closeModal}>
                <div className="modal-dialog modal-dialog-centered bd-example-modal-sm">
                    <div className="modal-content">
                        <div className="modal-header">
                            <h6 className="modal-title">{title}</h6>
                        </div>
                        <div className="modal-body">
                            <div className="row-container">
                                {context}
                            </div>
                        </div>
                        <div className="modal-footer">
                            <button type="button" dataClose="-1" className="btn btn-secondary">取消</button>
                            <button type="button" className="btn btn-primary" onClick={() => {
                                this.props.callback(this.props.value);
                                ReactDOM.unmountComponentAtNode(document.getElementById("modal"));
                            }}>确认</button>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

class ModalWindow extends React.Component {
    constructor(props) {
        super(props);
        const {title = "", itemName = "", value = ""} = this.props;
        this.state = {title, itemName, value};  // 初始化数据
        this.closeModal = this.closeModal.bind(this);
        this.inputChange = this.inputChange.bind(this);
        this.saveInfo = this.saveInfo.bind(this);
        this.blurActive = this.blurActive.bind(this);
        this.getFocus = this.getFocus.bind(this);
    }

    inputChange(event) {
        const v = event.currentTarget.value || "";
        this.setState({value: v, flag: "" != v});
    }

    closeModal(event) {
        const key = event.target.getAttribute("dataClose") || "";
        if (-1 == key) {
           ReactDOM.unmountComponentAtNode(document.getElementById("modal"));
        }
    }

    saveInfo() {
        const {value, flag} = this.state;
        this.props.callback(value, flag, this.props.value);
        ReactDOM.unmountComponentAtNode(document.getElementById("modal"));
    }

    blurActive(event) {
        const dom = event.currentTarget;
        dom.setAttribute("class", "input")
    }

    getFocus(event) {
        const dom = event.currentTarget;
        dom.setAttribute("class", "input isActive")
    }

    render() {
        const {title = "", itemName = "", value = "", flag} = this.state;
        return (
            <div class="modal fade show" tabindex="-1" dataClose="-1" onClick={this.closeModal}>
                <div class="modal-dialog modal-dialog-centered">
                    <div class="modal-content">
                    <div class="modal-header">
                        <h6 class="modal-title">{title}</h6>
                    </div>
                    <div class="modal-body">
                        <div class="label"><small>{itemName}</small></div>
                        <div class="row-container">
                            <div class="input-container">
                                <input onChange={this.inputChange} onFocus={this.getFocus} onBlur={this.blurActive} defaultValue={value} class="input"
                                autofocus="" tabindex="0" aria-label={itemName} />
                            </div>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="button" dataClose="-1" class="btn btn-secondary">取消</button>
                        {flag? <button type="button" className="btn btn-primary" onClick={this.saveInfo}>保存</button>:
                            <button type="button" className="btn btn-primary" disabled>保存</button>}
                    </div>
                    </div>
                </div>
            </div>
        )
    }
}

class MsgAlert extends React.Component {
    constructor(props) {
        super(props);
    }

    generateClass(type) {
        switch (type) {
            case "success": return "alert alert-success";
            case "error": return "alert alert-danger";
            default: return "alert alert-primary";
        }
    }

    componentDidMount() {
        const that = this;
        setTimeout(() => that.props.closeAlert(), 3* 1000)
    }

    render() {
        const {type, msg = "", closeAlert} = this.props;
        return (
            <div className={this.generateClass(type)} role="alert">
                {msg}
                <button type="button" className="close" onClick={closeAlert}>
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
        );
    }
}

class ModalWindow_editService extends React.Component {
    constructor(props) {
        super(props);
        const {title = "", serviceName = "", addr = ""} = this.props;
        this.state = {title, serviceName, addr};  // 初始化数据
        this.closeModal = this.closeModal.bind(this);
        this.inputChange = this.inputChange.bind(this);
        this.saveInfo = this.saveInfo.bind(this);
        this.blurActive = this.blurActive.bind(this);
        this.getFocus = this.getFocus.bind(this);
    }

    inputChange(event) {
        const dom = event.currentTarget;
        const k = dom.getAttribute("data-id");
        const v = dom.value || "";
        this.setState({[k]: v, flag: "" != v});
    }

    closeModal(event) {
        const key = event.target.getAttribute("dataClose") || "";
        if (-1 == key) {
            ReactDOM.unmountComponentAtNode(document.getElementById("modal"));
        }
    }

    saveInfo() {
        const {serviceName, addr, flag} = this.state;
        this.props.callback({serviceName, addr}, flag, {serviceName: this.props.serviceName, addr: this.props.addr});
        ReactDOM.unmountComponentAtNode(document.getElementById("modal"));
    }

    blurActive(event) {
        const dom = event.currentTarget;
        dom.setAttribute("class", "input")
    }

    getFocus(event) {
        const dom = event.currentTarget;
        dom.setAttribute("class", "input isActive")
    }

    render() {
        const {title = "", serviceName = "", addr = "", flag} = this.state;
        return (
            <div class="modal fade show" tabindex="-1" dataClose="-1" onClick={this.closeModal}>
                <div class="modal-dialog modal-dialog-centered">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h6 class="modal-title">{title}</h6>
                        </div>
                        <div class="modal-body">
                            <div class="label">
                                <small>服务名</small>
                            </div>
                            <div class="row-container">
                                <div class="input-container">
                                    <input onChange={this.inputChange} onFocus={this.getFocus} onBlur={this.blurActive} defaultValue={serviceName} class="input"
                                           data-id="serviceName" autofocus="" tabindex="0" aria-label={serviceName} />
                                </div>
                            </div>
                            <div className="label">
                                <small>服务地址</small>
                            </div>
                            <div className="row-container">
                                <div className="input-container">
                                    <input onChange={this.inputChange} onFocus={this.getFocus} onBlur={this.blurActive} defaultValue={addr} className="input"
                                           data-id="addr" tabIndex="0" aria-label={addr}/>
                                </div>
                            </div>
                        </div>
                        <div class="modal-footer">
                            <button type="button" dataClose="-1" class="btn btn-secondary">取消</button>
                            {flag? <button type="button" className="btn btn-primary" onClick={this.saveInfo}>保存</button>:
                                <button type="button" className="btn btn-primary" disabled>保存</button>}
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}