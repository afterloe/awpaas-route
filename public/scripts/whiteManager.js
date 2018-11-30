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

const jsonToQueryStr = data => {
    const item = [];
    Object.keys(data).map(key => item.push(`${key}=${encodeURIComponent(data[key])}`));
    return item.join("&");
};

const appendToRemote = (data, path) => new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.timeout = 15 * 1000;
    xhr.ontimeout = () => reject(new Error('time is up!'));
    xhr.open("PUT", path);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.setRequestHeader("cache-control", "no-cache");
    xhr.send(jsonToQueryStr(data));
    xhr.onreadystatechange = () => {
        if (4 === xhr.readyState) {
            const result = JSON.parse(xhr.responseText);
            if (200 === xhr.status) {
                if (200 !== result.code) reject(result.msg);
                resolve(result.data);
            } else reject(result.msg);
        }
    };
});

const deleteFromRemote = (data, path) => new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.timeout = 15 * 1000;
    xhr.ontimeout = () => reject(new Error('time is up!'));
    xhr.open("DELETE", path+"?"+jsonToQueryStr(data));
    xhr.setRequestHeader("cache-control", "no-cache");
    xhr.send();
    xhr.onreadystatechange = () => {
        if (4 === xhr.readyState) {
            const result = JSON.parse(xhr.responseText);
            if (200 === xhr.status) {
                if (200 !== result.code) reject(result.msg);
                resolve(result.data);
            } else reject(result.msg);
        }
    };
});

const getListFromRemote = path => new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.timeout = 15 * 1000;
    xhr.ontimeout = () => reject(new Error('time is up!'));
    xhr.open("get", path);
    xhr.send();
    xhr.onreadystatechange = () => {
        if (4 === xhr.readyState) {
            const result = JSON.parse(xhr.responseText);
            if (200 === xhr.status) {
                if (200 !== result.code) reject(result.msg);
                resolve(result.data);
            } else reject(result.msg);
        }
    };
});

class WhiteManager extends React.Component {
    constructor(props) {
        super(props);
        this.appendItem = this.appendItem.bind(this);
        this.editItem = this.editItem.bind(this);
        this.syncToRemote = this.syncToRemote.bind(this);
        this.renderMsgAlert = this.renderMsgAlert.bind(this);
        this.closeMsgAlert = this.closeMsgAlert.bind(this);
        this.appendItemToRemote = this.appendItemToRemote.bind(this);
        this.modifyToRemote = this.modifyToRemote.bind(this);
        this.deleteItem = this.deleteItem.bind(this);
        this.deleteToRemote = this.deleteToRemote.bind(this);
        this.state = {};
    }

    componentDidMount() {
        const that = this;
        getListFromRemote("manager/v1/whiteList").then(data => {
            that.setState({list: data}); // 初始化数据
        }).catch(error => {
            that.setState({list: [], msg: {type: "error", context: error}});
        });
    }

    deleteToRemote(data) {
        const that = this;
        const {msg = {}, list} = that.state;
        const index = list.findIndex(it => data === it);
        if (-1 !== index) {
            deleteFromRemote({item: data}, "manager/v1/whiteList").then(() => {
                list.splice(index, 1);
                Object.assign(msg, {type: "success", context: "删除成功..."});
                that.setState({msg, list});
            }).catch(error => {
                Object.assign(msg, {type: "error", context: error});
                that.setState({msg});
            });
        } else {
            Object.assign(msg, {type: "error", context: "删除失败..."});
            that.setState({msg});
        }

    }

    modifyToRemote(data, flag, oldData) {
        if (!flag) return;
        const that = this;
        const {msg = {}, list} = that.state;
        if (data === oldData) {
            Object.assign(msg, {type: "error", context: "未被修改..."});
            that.setState({msg});
            return;
        }
        const index = list.findIndex(it => oldData === it);
        if (-1 === index) {
            Object.assign(msg, {type: "error", context: "数据已被删除..."});
            that.setState({msg});
            return;
        }
        Promise.all([deleteFromRemote({item: oldData}, "manager/v1/whiteList")
            , appendToRemote({item: data}, "manager/v1/whiteList")]).then(() => {
            list[index] = data;
            Object.assign(msg, {type: "success", context: "修改成功..."});
            that.setState({msg, list});
        }).catch(error => {
            Object.assign(msg, {type: "error", context: error});
            that.setState({msg});
        });
    }

    appendItemToRemote(data, flag) {
        if (!flag) return ;
        const that = this;
        const {msg = {}, list} = that.state;
        const index = list.findIndex(it => data === it);
        if (-1 === index) {
            appendToRemote({item: data}, "manager/v1/whiteList").then(() => {
                list.push(data);
                Object.assign(msg, {type: "success", context: "保存成功..."});
                that.setState({msg, list});
            }).catch(error => {
                Object.assign(msg, {type: "error", context: error});
                this.setState({msg});
            });
        } else {
            Object.assign(msg, {type: "error", context: "失败：元素已存在..."});
            this.setState({msg});
        }

    }

    deleteItem(event) {
        const content = event.currentTarget.parentNode.previousSibling.textContent;
        const context = `确认删除 \t ${content} \t ?`;
        ReactDOM.render(<ModalWindow_alert title={"删除此项"} context={context} value={content} callback={this.deleteToRemote}/>, document.getElementById("modal"));
    }

    editItem(event) {
        const content = event.currentTarget.parentNode.previousSibling.textContent;
        ReactDOM.render(<ModalWindow title={"修改记录"} itemName={"白名单"} value={content} callback={this.modifyToRemote}/>, document.getElementById("modal"));
    }

    appendItem() {
        ReactDOM.render(<ModalWindow title={"添加记录"} itemName={"白名单"} callback={this.appendItemToRemote} />, document.getElementById("modal"));
    }

    renderMsgAlert(msg) {
        const {type, context = ""} = msg;
        return "" === context? "": <MsgAlert type= {type} msg= {context} closeAlert= {this.closeMsgAlert}/>
    }

    closeMsgAlert() {
        this.setState({msg: {}})
    }

    syncToRemote() {
        const that = this;
        that.setState({msg: {type: "info", context: "同步中..."}});
        getListFromRemote("manager/v1/whiteList").then(data => {
            that.setState({list: data}); // 初始化数据
        }).catch(error => {
            that.setState({list: [], msg: {type: "error", context: error}});
        });
    }

    renderList(list = []) {
        return list.map(it => (
            <div className="media text-muted pt-3">
                <div className="media-body pb-3 mb-0 small lh-125 border-bottom border-gray">
                    <div className="d-flex justify-content-between align-items-center w-100">
                        <strong className="text-gray-dark">{it}</strong>
                        <span>
                            <span className="cont-btn" onClick={this.editItem}>
                                <embed src="images/edit.svg" width="16" height="16" type="image/svg+xml"/>
                                <span>修改</span>
                            </span>
                            <span className="cont-btn" onClick={this.deleteItem}>
                                <embed src="images/trash.svg" width="16" height="16" type="image/svg+xml"/>
                                <span>删除</span>
                            </span>
                        </span>
                    </div>
                </div>
            </div>
        ));
    }

    render() {
        const {msg = {}, list = []} = this.state;
        return (
            <main role="main" className="col-md-9 ml-sm-auto col-lg-10 px-4">
                <div
                    className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
                    <h1 className="h2">网关白名单</h1>
                </div>
                {this.renderMsgAlert(msg)}
                <div class="my-3 p-3 bg-white rounded shadow-sm m-cent">
                    <h6 class="border-bottom border-gray pb-2 mb-0">拦截信息</h6>
                    <div class="media text-muted pt-3">
                        <div class="media-body pb-3 mb-0 small lh-125 border-bottom border-gray">
                            <div class="d-flex justify-content-between align-items-center w-100">
                                <strong class="text-gray-dark">请求头</strong>
                                <span class="cont-btn">
                                    <embed src="images/settings.svg" width="16" height="16" type="image/svg+xml"/>
                                    <span>设置</span>
                                </span>
                            </div>
                            <span class="d-block">access-token</span>
                        </div>
                    </div>
                    <div class="media text-muted pt-3">
                        <div class="media-body pb-3 mb-0 small lh-125 border-bottom border-gray">
                            <div class="d-flex justify-content-between align-items-center w-100">
                                <strong class="text-gray-dark">启用拦截</strong>
                                <span class="cont-btn">
                                    <span>ON</span>
                                </span>
                            </div>
                            <span class="d-block">未携带请求头的链接进行自动拦截</span>
                        </div>
                    </div>
                    <small class="d-block text-right mt-3 cont-btn" onClick={this.syncToRemote}>
                        <embed src="images/refresh-cw.svg" width="16" height="16" type="image/svg+xml"/>
                        <span>同步</span>
                    </small>
                </div>
                <div class="my-3 p-3 bg-white rounded shadow-sm m-cent">
                    <h6 class="border-bottom d-flex justify-content-between align-items-center">
                        <span class="d-block">列表</span>
                        <small class="d-block text-right mt-3 mb-3">
                            <span class="cont-btn" onClick={this.syncToRemote}>
                                <embed src="images/refresh-cw.svg" width="16" height="16" type="image/svg+xml"/>
                                <span>同步</span>
                            </span>
                            <span class="cont-btn" onClick={this.appendItem}>
                                <embed src="images/plus-circle.svg" width="16" height="16" type="image/svg+xml"/>
                                <span>添加</span>
                            </span>
                        </small>
                    </h6>
                    {this.renderList(list)}
                </div>
                <div id="modal"></div>
            </main>
        );
    }
}