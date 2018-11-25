"use strict";

class Header extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <nav class="navbar navbar-dark fixed-top bg-dark flex-md-nowrap p-0 shadow">
                <a class="navbar-brand col-sm-3 col-md-2 mr-0" href="#">前置数据网关</a>
                <input class="form-control form-control-dark w-100" type="text" placeholder="搜索" aria-label="搜索">
                    <ul class="navbar-nav px-3">
                        <li class="nav-item text-nowrap">
                            <a class="nav-link" href="#">登出</a>
                        </li>
                    </ul>
            </nav>
        )
    }
}
        
ReactDOM.render(<Header />, document.getElementById("head")