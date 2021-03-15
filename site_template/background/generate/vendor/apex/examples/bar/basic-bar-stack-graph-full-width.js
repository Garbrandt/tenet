var options = {
	chart: {
		height: 300,
		type: 'bar',
		stacked: true,
		stackType: '100%',
		toolbar: {
      show: false,
    },
    dropShadow: {
			enabled: true,
			opacity: 0.1,
			blur: 5,
			left: -10,
			top: 10
		},
	},
	plotOptions: {
		bar: {
			horizontal: true,
		},
	},
	stroke: {
		show: false,
	},
	series: [{
		name: 'iPhone',
		data: [44, 55, 41, 37, 22, 43, 21]
	},{
		name: 'iPad',
		data: [53, 32, 33, 52, 13, 43, 32]
	},{
		name: 'iMac',
		data: [12, 17, 11, 9, 15, 11, 20]
	},{
		name: 'MacBook',
		data: [9, 7, 5, 8, 6, 9, 4]
	},{
		name: 'MacMini',
		data: [25, 12, 19, 32, 25, 24, 10]
	}],
	xaxis: {
		categories: [2008, 2009, 2010, 2011, 2012, 2013, 2014],
	},
	tooltip: {
		y: {
			formatter: function(val) {
				return val + "K"
			}
		}
	},
	fill: {
		opacity: 1
	},
	legend: {
		position: 'bottom',
		horizontalAlign: 'center',
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
	colors: ['#007ae1', '#ff3e3e', '#00bb42', '#ffbf05'],
}

var chart = new ApexCharts(
	document.querySelector("#basic-bar-stack-graph-full-width"),
	options
);

chart.render();
