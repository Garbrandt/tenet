var options = {
	chart: {
		height: 300,
		type: 'area',
		toolbar: {
			show: false,
		},
		dropShadow: {
			enabled: true,
			opacity: 0.3,
			blur: 5,
			left: -10,
			top: 20
		},
	},
	colors: ['#007ae1', '#ff3e3e'],
	dataLabels: {
		enabled: false,
	},
	legend: {
  	show: false,
  },
	stroke: {
		show: true,
		curve: 'smooth',
		width: 3,
		lineCap: 'square'
	},
	series: [{
		name: 'Revenue',
		data: [5000, 40000, 50000, 90000, 100000, 85000, 95000]
	},{
    name: 'Sales',
    data: [2500, 18000, 22000, 43000, 81000, 72000, 37000]
  }],
	xaxis: {
		axisBorder: {
			show: false
		},
		categories: ["SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"],
		axisTicks: {
			show: true
		},
		crosshairs: {
			show: true
		},
		labels: {
			offsetX: 0,
			offsetY: 5,
		}
	},
	yaxis: {
		labels: {
			formatter: function(value, index) {
				return (value / 1000) + 'K'
			},
			offsetX: -15,
			offsetY: 20,
		}
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
			right: 0,
			bottom: 0,
			left: 0
		}, 
	}, 
	fill: {
		type:"gradient",
		gradient: {
			type: "vertical",
			shadeIntensity: 1,
			inverseColors: !1,
			opacityFrom: .3,
			opacityTo: .05,
			stops: [15, 100]
		}
	},
}

var chart = new ApexCharts(
	document.querySelector("#lineRevenueGraph"),
	options
);

chart.render();