"use strict";

class WhiteManager extends React.Component {
    constructor(props) {
        super(props);
        this.state = {}; // 初始化数据
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
                                <span>添加</span>
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
            </main>
        );
    }
}