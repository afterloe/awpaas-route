"use strict";

const systemMenu = [{
    name: "首页",
    icon: "images/home.svg",
    isClick: true,
    index: "main"
}, {
    name: "白名单管理",
    icon: "images/file.svg",
    index: "whiteManager"
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

const requestTotal = [23, 233, 122, 309, 177, 133, 12];
const cpuTotal = [3, 12, 20, 12, 45, 124, 18];
const reqRank = [{name: "docker", url: "user/images/json", count: "50002", trend: "-"},
    {name: "docker", url: "user/images/json", count: "50002", trend: "-"},
    {name: "docker", url: "user/images/json", count: "50002", trend: "UP"},
    {name: "docker", url: "user/images/json", count: "50002", trend: "DOWN"},
    {name: "docker", url: "user/images/json", count: "50002", trend: "DOWN"},
    {name: "docker", url: "user/images/json", count: "50002", trend: "-"}];

class GateWay extends React.Component {
    constructor(props) {
        super(props);
        const {menu = [], reqRank = [], reqTotal = [], cpuTotal = []}= props;
        this.state = {menu, reqRank, reqTotal, cpuTotal};
        this.clickItem = this.clickItem.bind(this);
    }

    clickItem(event) {
        const key = event.currentTarget.getAttribute("data-index") || "";
        if ("" === key) return;
        this.setState(prevState => {
            const {menu} = prevState;
            return {menu: menu.map(it => {
                it.isClick = key === it.index? true: false;
                return it;
            })}
        });
    }

    render() {
        const {menu, reqRank, reqTotal, cpuTotal} = this.state;
        return (
            <div class="row">
                <NavLeft menu={menu} links={this.props.links} clickItem={this.clickItem}/>
                <TotalMain rank={reqRank} reqTotal={reqTotal} cpuTotal={cpuTotal}/>
            </div>
        )
    }
}

ReactDOM.render(<Header name="前置数据网关" version="v1.0.3"/>, document.getElementById("head"));
ReactDOM.render(<GateWay menu={systemMenu} links={linkMenu} reqRank={reqRank} reqTotal={requestTotal}
                         cpuTotal={cpuTotal}/>, document.getElementById("app"));