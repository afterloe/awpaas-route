"use strict";

class NavLeft extends React.Component {
    constructor(props) {
        super(props);
        this.state = props;
        this.clickItem = this.clickItem.bind(this);
    }

    clickItem(event) {
        const key = event.currentTarget.getAttribute("data-index") || "";
        // const k = key.getAgetAttribute("data-index") || "";
        console.log(key)
        // console.log(k)
        if ("" === key) return;
    }

    renderMenu() {
        const {menu = []} = this.state;
        return menu.map(it => (
            <li class="nav-item" onClick={this.clickItem} data-index={it.index}>
                <a class={it.isClick? "nav-link active":"nav-link"} href="#">
                    <embed src={it.icon} width="16" height="16" type="image/svg+xml"/>
                    {it.name} {it.isClick? (<span class="sr-only">(current)</span>): ""}
                </a>
            </li>
        ))
    }

    renderLink() {
        const {links = []} = this.state;
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

