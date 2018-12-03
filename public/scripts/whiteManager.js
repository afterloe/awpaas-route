"use strict";

const generatorModify= async (old, n, path) => {
    await deleteFromRemote(old, path);
    await appendToRemote(n, path);
};

const whiteManagerURL = "manager/v1/whiteList";

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
        getListFromRemote(whiteManagerURL).then(data => {
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
            deleteFromRemote({item: data}, whiteManagerURL).then(() => {
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
        generatorModify({item: oldData}, {item: data}, whiteManagerURL).then(() => {
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
            appendToRemote({item: data}, whiteManagerURL).then(() => {
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
        getListFromRemote(whiteManagerURL).then(data => {
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