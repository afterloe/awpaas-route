"use strict";

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
                                           data-id="addr" autoFocus="" tabIndex="0" aria-label={addr}/>
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

class ServiceRegistry extends React.Component {
    constructor(props) {
        super(props);
        this.showMenu = this.showMenu.bind(this);
        this.itemClick = this.itemClick.bind(this);
        this.openModal = this.openModal.bind(this);
        this.appendItemToRemote = this.appendItemToRemote.bind(this);
        this.renderServiceMap = this.renderServiceMap.bind(this);
        this.renderMsgAlert = this.renderMsgAlert.bind(this);
        this.openDelWindow = this.openDelWindow.bind(this);
        this.deleteToRemote = this.deleteToRemote.bind(this);
        this.openModifyWindow = this.openModifyWindow.bind(this);
        this.modifyItemToRemote = this.modifyItemToRemote.bind(this);
        this.state = {serviceList: [{serviceName: "couchdb", addr: "127.0.0.1:8088"}]};
    }

    static closeMenu(dom) {
        const flag = dom.getAttribute("data-flag");
        if ("f" === flag) {
            dom.setAttribute("data-flag", "t");
            dom.setAttribute("class", "dropdown-menu show")
        } else {
            dom.setAttribute("data-flag", "f");
            dom.setAttribute("class", "dropdown-menu")
        }
    }

    modifyItemToRemote(data, flag) {
        const [
            {serviceList= [], msg= {}},
            {serviceName}
        ] = [
            this.state,
            data
        ];
        if (!flag) {
            return ;
        }
        const index = serviceList.findIndex(item => {
            return item.serviceName == serviceName;
        });
        if (-1 === index) {
            Object.assign(msg, {type: "error", context: "服务映射不存在于注册中心..."});
            this.setState({msg});
            return 
        }
        serviceList[index] = data;
        Object.assign(msg, {type: "success", context: "服务修改成功..."});
        this.setState({serviceList, msg});
    }

    appendItemToRemote(data, flag) {
        const [
            {serviceList= [], msg= {}},
            {serviceName, addr}
        ] = [
            this.state,
            data
        ];
        const index = serviceList.findIndex(item => {
            return item.serviceName == serviceName;
        });
        if (-1 !== index) {
            Object.assign(msg, {type: "error", context: "服务已存在..."});
            this.setState({msg});
            return 
        }
        serviceList.push({serviceName, addr});
        Object.assign(msg, {type: "success", context: "服务注册成功..."});
        this.setState({serviceList, msg});
    }

    openModal() {
        ReactDOM.render(<ModalWindow_editService title={"添加服务注册信息"} callback={this.appendItemToRemote} />, document.getElementById("modal"));
    }

    deleteToRemote(index) {
        const {serviceList= [], msg= {}} = this.state;
        const service = serviceList[index];
        Object.assign(msg, {type: "success", context: "删除成功..."});
        if (-1 !== index) {
            serviceList.splice(index, 1)
        }
        this.setState({msg, serviceList});
    }

    openModifyWindow(index) {
        const {serviceList = []} = this.state;
        const service = serviceList[index];
        ReactDOM.render(<ModalWindow_editService title={"修改服务注册信息"} serviceName={service.serviceName} addr={service.addr} callback={this.modifyItemToRemote} />, document.getElementById("modal"));
    }

    openDelWindow(index) {
        const {serviceList = []} = this.state;
        const service = serviceList[index];
        const context = `确认删除 \t ${service.serviceName} \t ? \r\n 映射地址为 ${service.addr}`;
        ReactDOM.render(<ModalWindow_alert title={"删除此项"} context={context} value={index} callback={this.deleteToRemote}/>, document.getElementById("modal"));
    }

    itemClick(event) {
        const item = event.target;
        const [
            cmd,
            index,
        ]= [
            item.getAttribute("data-cmd") || "", 
            item.parentNode.getAttribute("data-index"),
        ];
        if ("" === cmd) {
            ServiceRegistry.closeMenu(item)
        }
        ServiceRegistry.closeMenu(item.parentNode);
        switch (cmd) {
            case "del": 
                this.openDelWindow(index);
                return;
            case "modify":
                this.openModifyWindow(index);
                return;
            default:
                return;
        }
    }

    renderMsgAlert() {
        const {msg = {}} = this.state;
        const {type, context = ""} = msg;
        return "" === context? "": <MsgAlert type= {type} msg= {context} closeAlert= {()=>this.setState({msg: {}})}/>
    }

    showMenu(event) {
        const menu = event.currentTarget.nextSibling;
        ServiceRegistry.closeMenu(menu);
    }

    renderServiceMap() {
        const {serviceList = []} = this.state;
        return 0 === serviceList.length? ( 
        <div className="media text-muted pt-3">
            暂无数据!
        </div>): serviceList.map((item, index) => (
            <div className="media text-muted pt-3">
                <div className="media-body pb-3 mb-0 small lh-125 border-bottom border-gray">
                    <div className="d-flex justify-content-between align-items-center w-100">
                        <strong className="text-gray-dark">
                            <span className="badge badge-success">正常</span> {item.serviceName}
                        </strong>
                        <div className="cont-btn btn-group show" role="group">
                            <svg onClick={this.showMenu} xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                                viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" 
                                stroke-linecap="round" stroke-linejoin="round" className="feather feather-more-vertical">
                                <circle cx="12" cy="12" r="1"></circle>
                                <circle cx="12" cy="5" r="1"></circle>
                                <circle cx="12" cy="19" r="1"></circle>
                            </svg>
                            <div className="dropdown-menu" data-index={index} data-flag="f" onClick={this.itemClick}>
                                <span className="dropdown-item" data-cmd="del">删除</span>
                                <span className="dropdown-item" data-cmd="modify">修改</span>
                                <span className="dropdown-item" data-cmd="detail">详情</span>
                                <span className="dropdown-item" data-cmd="rely">依赖关系</span>
                            </div>
                        </div>
                    </div>
                    <span className="d-block detail">
                        <span>map to: <strong>{item.addr}</strong></span>
                        <span>register time: <strong>{item.registryTime || new Date().toLocaleString()}</strong></span>
                    </span>
                </div>
            </div>
        ));
    }

    render() {
        return (
            <main role="main" className="col-md-9 ml-sm-auto col-lg-10 px-4">
                <div
                    className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
                    <h1 className="h2">服务注册列表</h1>
                    <div className="btn-toolbar mb-2 mb-md-0">
                        <div className="btn-group mr-2">
                            <button className="btn btn-sm btn-outline-secondary cont-btn" onClick={this.openModal}>
                                <embed src="images/upload-cloud.svg" width="16" height="16" type="image/svg+xml"/>
                                <span>手动注册</span>
                            </button>
                            <button className="btn btn-sm btn-outline-secondary cont-btn">
                                <embed src="images/aperture.svg" width="16" height="16" type="image/svg+xml"/>
                                <span>扫描</span>
                            </button>
                        </div>
                    </div>
                </div>
                <div className="my-3 p-3 bg-white rounded shadow-sm m-cent">
                    <h6 className="border-bottom d-flex justify-content-between align-items-center">
                    </h6>
                        {this.renderMsgAlert()}
                        {this.renderServiceMap()}
                    <div id="modal"></div>
                </div>
            </main>
        );
    }
}