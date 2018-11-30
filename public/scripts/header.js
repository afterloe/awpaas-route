"use strict";
/**
 *  顶部导航栏
 */
class Header extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        let p = this.props;
        return (
            <nav class="navbar navbar-dark fixed-top bg-dark flex-md-nowrap p-0 shadow">
            <a class="navbar-brand col-sm-3 col-md-2 mr-0" href="/#">{p.name}</a>
        <input class="form-control form-control-dark w-100" type="text" placeholder="搜索" aria-label="搜索" />
            <ul class="navbar-nav px-3">
            <li class="nav-item text-nowrap">
            <a class="nav-link" href="#">{p.version}</a>
        </li>
        </ul>
        </nav>
    )
    }
}
