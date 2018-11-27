"use strict";

const labels = ["周天", "周一", "周二", "周三", "周四", "周五", "周六"];

class TotalMain extends React.Component {
    constructor(props) {
        super(props);
    }

    componentDidMount() {
        const {reqTotal = [], cpuTotal = []} = this.props; // 拉取数据
        TotalMain.loadChart(reqTotal, cpuTotal) // 绘制流量报表
    }

    renderRequestRank() {
        const {rank = []} = this.props;
        return rank.map((it, i) => (
            <tr>
                <td>{i}</td>
                <td>{it.name}</td>
                <td>{it.url}</td>
                <td>{it.count}</td>
                <td>{it.trend}</td>
            </tr>
        ));
    }

    static loadChart(requestTotal, cpuTotal) {
        new Chart(document.getElementById("myChart"), {
            type: 'line',
            data: {
                labels,
                datasets: [{
                    data: requestTotal,
                    lineTension: 0,
                    backgroundColor: 'transparent',
                    borderColor: '#f94d00', // 线条颜色
                    borderWidth: 4,
                    pointBackgroundColor: '#f94d00' // 提示颜色
                }, {
                    data: cpuTotal,
                    lineTension: 0,
                    backgroundColor: 'transparent',
                    borderColor: '#4f86f7', // 线条颜色
                    borderWidth: 4,
                    pointBackgroundColor: '#4f86f7' // 提示颜色
                }]
            },
            options: {
                scales: {
                    yAxes: [{
                        ticks: {
                            beginAtZero: false
                        }
                    }]
                },
                legend: {
                    display: false,
                }
            }
        });
    }

    render() {
        return (
            <main role="main" class="col-md-9 ml-sm-auto col-lg-10 px-4">
                <div class="chartjs-size-monitor">
                    <div class="chartjs-size-monitor-expand">
                        <div></div>
                    </div>
                    <div class="chartjs-size-monitor-shrink">
                        <div></div>
                    </div>
                </div>
                <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
                    <h1 class="h2">网关流量</h1>
                    <div class="btn-toolbar mb-2 mb-md-0">
                        <div class="btn-group mr-2">
                            <button class="btn btn-sm btn-outline-secondary">分享</button>
                            <button class="btn btn-sm btn-outline-secondary">打印</button>
                        </div>
                        <button class="btn btn-sm btn-outline-secondary dropdown-toggle">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor"
                                 stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-calendar">
                                <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
                                <line x1="16" y1="2" x2="16" y2="6"></line>
                                <line x1="8" y1="2" x2="8" y2="6"></line>
                                <line x1="3" y1="10" x2="21" y2="10"></line>
                            </svg>
                            本周
                        </button>
                    </div>
                </div>

                <canvas class="my-4 w-100 chartjs-render-monitor" id="myChart" width="1053" height="444"></canvas>

                <h2>访问排行</h2>
                <div class="table-responsive">
                    <table class="table table-striped table-sm">
                        <thead>
                        <tr>
                            <th>#</th>
                            <th>服务名</th>
                            <th>URL</th>
                            <th>访问次数</th>
                            <th>趋势</th>
                        </tr>
                        </thead>
                        <tbody>{this.renderRequestRank()}</tbody>
                    </table>
                </div>
            </main>
        )
    }
}