var options = {
	chart: {
		height: 350,
		type: 'line',
		zoom: {
			enabled: false
		},
		dropShadow: {
			enabled: true,
			opacity: 0.1,
			blur: 5,
			left: -10,
			top: 10
		},
	},
	dataLabels: {
		enabled: false
	},
	colors: ['#007ae1', '#ff3e3e', '#00bb42', '#ffbf05'],
	stroke: {
		width: [3, 3, 3],
		curve: 'straight',
		dashArray: [0, 8, 5]
	},
	series: [
		{
			name: "Session Duration",
			data: [45, 52, 38, 24, 33, 26, 21, 20, 6, 8, 15, 10]
		},
		{
			name: "Page Views",
			data: [35, 41, 62, 42, 13, 18, 29, 37, 36, 51, 32, 35]
		},
		{
			name: 'Total Visits',
			data: [87, 57, 74, 99, 75, 38, 62, 47, 82, 56, 45, 47]
		}
	],
	title: {
		text: 'Page Statistics',
		align: 'center'
	},
	markers: {
		size: 0,
		hover: {
			sizeOffset: 6
		}
	},
	xaxis: {
		categories: ['01 Jan', '02 Jan', '03 Jan', '04 Jan', '05 Jan', '06 Jan', '07 Jan', '08 Jan', '09 Jan',
			'10 Jan', '11 Jan', '12 Jan'
		],
	},
	tooltip: {
		y: [{
			title: {
				formatter: function (val) {
					return val + " (mins)"
				}
			}
		}, {
			title: {
				formatter: function (val) {
					return val + " per session"
				}
			}
		}, {
			title: {
				formatter: function (val) {
					return val;
				}
			}
		}]
	},
	grid: {
    borderColor: '#ced1e0',
    strokeDashArray: 5,
    xaxis: {
      lines: {
        show: true
      }
    },   
    yaxis: {
      lines: {
        show: false,
      } 
    },
    padding: {
      top: 0,
      right: 20,
      bottom: 0,
      left: 30
    }, 
  },
}

var chart = new ApexCharts(
	document.querySelector("#stepLineChart"),
	options
);

chart.render();