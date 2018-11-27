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
        return (
            <div class="modal fade show" tabindex="-1" dataClose="-1" onClick={this.closeModal}>
                <div class="modal-dialog modal-dialog-centered">
                    <div class="modal-content">
                    <div class="modal-header">
                        <h6 class="modal-title">添加记录</h6>
                    </div>
                    <div class="modal-body">
                        <div class="label"><small>白名单</small></div>
                        <div class="row-container">
                            <div class="input-container">
                                <input class="input" autofocus="" tabindex="0" aria-label="Search engine" />
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

class WhiteManager extends React.Component {
    constructor(props) {
        super(props);
        this.appendItem = this.appendItem.bind(this);
        this.state = {}; // 初始化数据
    }

    appendItem(event) {
        ReactDOM.render(<AppendItem />, document.getElementById("modal"));
    }

    render() {
        return (
            <main role="main" className="col-md-9 ml-sm-auto col-lg-10 px-4">
                <div
                    className="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
                    <h1 className="h2">网关白名单</h1>
                </div>
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
                    <small class="d-block text-right mt-3 cont-btn">
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
                            <span class="cont-btn">
                                <embed src="images/plus-circle.svg" width="16" height="16" type="image/svg+xml"/>
                                <span onClick={this.appendItem}>添加</span>
                            </span>
                        </small>
                    </h6>
                    <div class="media text-muted pt-3">
                        <div class="media-body pb-3 mb-0 small lh-125 border-bottom border-gray">
                            <div class="d-flex justify-content-between align-items-center w-100">
                                <strong class="text-gray-dark">/member/login</strong>
                                <span>
                                    <span class="cont-btn">
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