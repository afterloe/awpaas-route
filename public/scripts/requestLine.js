feather.replace()
var ctx = document.getElementById("myChart");
var myChart = new Chart(ctx, {
    type: 'line',
    data: {
        labels: ["周天", "周一", "周二", "周三", "周四", "周五", "周六"],
        datasets: [{
            data: [23, 233, 122, 309, 177, 133, 12],
            lineTension: 0,
            backgroundColor: 'transparent',
            borderColor: '#f94d00', // 线条颜色
            borderWidth: 4,
            pointBackgroundColor: '#f94d00' // 提示颜色
        }, {
            data: [123, 333, 222, 409, 277, 233, 112],
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