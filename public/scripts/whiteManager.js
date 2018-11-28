"use strict";

class AppendItem extends React.Component {
    constructor(props) {
        super(props);
        // this.state = {}; // 初始化数据
        this.closeModal = this.closeModal.bind(this);
    }

    closeModal(event) {
        const key = event.target.getAttribute("dataClose") || "";
        if (-1 == key) {
           ReactDOM.unmountComponentAtNode(document.getElementById("modal"))
        }
    }

    render() {
        const {title = "", itemName = "", value = ""} = this.props;
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
                                <input class="input" value={value} autofocus="" tabindex="0" aria-label={itemName} />
                                <div class="underline"></div>
                            </div>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="button" dataClose="-1" class="btn btn-secondary">取消</button>
                        <button type="button" class="btn btn-primary" disabled>保存</button>
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

class WhiteManager extends React.Component {
    constructor(props) {
        super(props);
        this.appendItem = this.appendItem.bind(this);
        this.editItem = this.editItem.bind(this);
        this.syncToRemote = this.syncToRemote.bind(this);
        this.renderMsgAlert = this.renderMsgAlert.bind(this);
        this.closeMsgAlert = this.closeMsgAlert.bind(this);
        this.state = {}; // 初始化数据
    }

    editItem(event) {
        const content = event.currentTarget.parentNode.previousSibling.textContent;
        ReactDOM.render(<AppendItem title={"修改记录"} itemName={"白名单"} value={content}/>, document.getElementById("modal"));
    }

    appendItem() {
        ReactDOM.render(<AppendItem title={"添加记录"} itemName={"白名单"} />, document.getElementById("modal"));
    }

    renderMsgAlert(msg) {
        const {type, context = ""} = msg;
        return "" === context? "": <MsgAlert type= {type} msg= {context} closeAlert= {this.closeMsgAlert}/>
    }

    closeMsgAlert() {
        this.setState({msg: {}})
    }

    syncToRemote() {
        this.setState({msg: {type: "info", context: "同步中..."}})
    }

    render() {
        const {msg = {}} = this.state;
        return (
            <main role="main" className="col-md-9 ml-sm-auto col-lg-10 px-4">
                <div
                    className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
                    <h1 className="h2">网关白名单</h1>
                </div>
                {this.renderMsgAlert(msg)}
                <div class="my-3 p-3 bg-white rounded shadow-sm">
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
                <div class="my-3 p-3 bg-white rounded shadow-sm">
                    <h6 class="border-bottom d-flex justify-content-between align-items-center">
                        <span class="d-block">列表</span>
                        <small class="d-block text-right mt-3 mb-3">
                            <span class="cont-btn">
                                <embed src="images/refresh-cw.svg" width="16" height="16" type="image/svg+xml"/>
                                <span>同步</span>
                            </span>
                            <span class="cont-btn" onClick={this.appendItem}>
                                <embed src="images/plus-circle.svg" width="16" height="16" type="image/svg+xml"/>
                                <span>添加</span>
                            </span>
                        </small>
                    </h6>
                    <div class="media text-muted pt-3">
                        <div class="media-body pb-3 mb-0 small lh-125 border-bottom border-gray">
                            <div class="d-flex justify-content-between align-items-center w-100">
                                <strong class="text-gray-dark">/member/login</strong>
                                <span>
                                    <span class="cont-btn" onClick={this.editItem}>
                                        <embed src="images/edit.svg" width="16" height="16" type="image/svg+xml"/>
                                        <span>修改</span>
                                    </span>
                                    <span class="cont-btn">
                                        <embed src="images/trash.svg" width="16" height="16" type="image/svg+xml"/>
                                        <span>删除</span>
                                    </span>
                                </span>
                            </div>
                       </div>
                    </div>
                </div>
                <div id="modal"></div>
            </main>
        );
    }
}