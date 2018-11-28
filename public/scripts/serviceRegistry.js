"use strict";

class ModalWindow_editService extends React.Component {
    constructor(props) {
        super(props);
        const {title = "", serviceName = "", addr = ""} = this.props;
        this.state = {title, serviceName, addr};  // 初始化数据
        this.closeModal = this.closeModal.bind(this);
        this.inputChange = this.inputChange.bind(this);
        this.saveInfo = this.saveInfo.bind(this);
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
                                    <input onChange={this.inputChange} defaultValue={serviceName} class="input"
                                           data-id="serviceName" autofocus="" tabindex="0" aria-label={serviceName} />
                                    <div class="underline"></div>
                                </div>
                            </div>
                            <div className="label">
                                <small>服务地址</small>
                            </div>
                            <div className="row-container">
                                <div className="input-container">
                                    <input onChange={this.inputChange} defaultValue={addr} className="input"
                                           data-id="addr" autoFocus="" tabIndex="0" aria-label={addr}/>
                                    <div className="underline"></div>
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

    appendItemToRemote(data, flag) {
        console.log(data, flag)
    }

    openModal() {
        ReactDOM.render(<ModalWindow_editService title={"添加服务映射"} callback={this.appendItemToRemote} />, document.getElementById("modal"));
    }

    itemClick(event) {
        const item = event.target;
        const cmd = item.getAttribute("data-cmd") || "";
        if ("" === cmd) {
            ServiceRegistry.closeMenu(item)
        }
        console.log(cmd);
        ServiceRegistry.closeMenu(item.parentNode);
    }

    showMenu(event) {
        const menu = event.currentTarget.nextSibling;
        ServiceRegistry.closeMenu(menu);
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
                    <div className="media text-muted pt-3">
                            <div className="media-body pb-3 mb-0 small lh-125 border-bottom border-gray">
                                <div className="d-flex justify-content-between align-items-center w-100">
                                    <strong className="text-gray-dark">
                                        <span className="badge badge-success">正常</span> couchdb
                                    </strong>
                                    <div className="cont-btn btn-group show" role="group">
                                        <svg onClick={this.showMenu} xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                                             viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
                                             stroke-linecap="round" stroke-linejoin="round"
                                             className="feather feather-more-vertical">
                                            <circle cx="12" cy="12" r="1"></circle>
                                            <circle cx="12" cy="5" r="1"></circle>
                                            <circle cx="12" cy="19" r="1"></circle>
                                        </svg>
                                        <div className="dropdown-menu" data-flag="f" onClick={this.itemClick}>
                                            <span className="dropdown-item" data-cmd="del">删除</span>
                                            <span className="dropdown-item" data-cmd="modify">修改</span>
                                            <span className="dropdown-item" data-cmd="detail">详情</span>
                                            <span className="dropdown-item" data-cmd="rely">依赖关系</span>
                                        </div>
                                    </div>
                                </div>
                                <span className="d-block detail">
                                    <span>map to: 127.0.0.1:8088</span>
                                    <span>register time: 2018-11-28 14:24:33</span>
                                </span>
                            </div>
                    </div>
                    <div id="modal"></div>
                </div>
            </main>
        );
    }
}