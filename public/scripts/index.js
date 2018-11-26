"use strict";

const systemMenu = [{
    name: "首页",
    icon: "images/home.svg",
    isClick: true,
    index: "main"
}, {
    name: "白名单管理",
    icon: "images/file.svg",
    index: "whilteManager"
}, {
    name: "服务注册列表",
    icon: "images/layers.svg",
    index: "serviceRegistry"
}, {
    name: "活跃用户",
    icon: "images/users.svg",
    index: "busyUsers"
}, {
    name: "网关状态",
    icon: "images/bar-chart-2.svg",
    index: "gatewayStatus"
}];
const linkMenu = [{
    name: "统一管理子系统",
    href: "https://127.0.0.1:8088"
}, {
    name: "统一认证及审计子系统",
    href: "https://127.0.0.1:8088"
}, {
    name: "蜂窝式数据仓库管理子系统",
    href: "https://127.0.0.1:8088"
}, {
    name: "数据采集清洗引擎",
    href: "https://127.0.0.1:8088"
}];

class GateWay extends React.Component {
    constructor(props) {
        super(props);
        const {menu} = props;
        this.state = {menu};
        this.clickItem = this.clickItem.bind(this);
    }

    clickItem(event) {
        const key = event.currentTarget.getAttribute("data-index") || "";
        if ("" === key) return;
        // TODO 切换
        alert(key)
    }

    render() {
        return (
            <div class="row">
                <NavLeft menu={this.state.menu} links={this.props.links} clickItem={this.clickItem}/>
                <TotalMain />
            </div>
        )
    }
}

ReactDOM.render(<Header name="前置数据网关" version="v1.0.3"/>, document.getElementById("head"));
ReactDOM.render(<GateWay menu={systemMenu} links={linkMenu}/>, document.getElementById("app"))