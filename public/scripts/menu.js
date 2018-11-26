"use strict";

class NavLeft extends React.Component {
    constructor(props) {
        super(props);
    }

    renderMenu() {
        const {menu = []} = this.props;
        return menu.map(it => (
            <li class="nav-item" onClick={this.props.clickItem} data-index={it.index}>
                <a class={it.isClick? "nav-link active":"nav-link"} href="#">
                    <embed src={it.icon} width="16" height="16" type="image/svg+xml"/>
                    {it.name} {it.isClick? (<span class="sr-only">(current)</span>): ""}
                </a>
            </li>
        ))
    }

    renderLink() {
        const {links = []} = this.props;
        return links.map(it => (
            <li class="nav-item">
                <a class="nav-link" href={it.href}>
                    <embed src="images/link.svg" width="16" height="16" type="image/svg+xml"/>
                    {it.name}
                </a>
            </li>
        ))
    }

    render() {
        return (
            <nav class="col-md-2 d-none d-md-block bg-light sidebar">
                <div class="sidebar-sticky">
                    <ul class="nav flex-column">
                        {this.renderMenu()}
                    </ul>

                    <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted">
                        <span>常用链接</span>
                        <a class="d-flex align-items-center text-muted link" href="#">
                            <embed src="images/plus-circle.svg" width="16" height="16" type="image/svg+xml"/>
                        </a>
                    </h6>
                    <ul class="nav flex-column mb-2">
                        {this.renderLink()}
                    </ul>
                </div>
            </nav>
        )
    }
}

