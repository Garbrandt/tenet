// Calendar List View
document.addEventListener('DOMContentLoaded', function() {
	var calendarEl = document.getElementById('calendarListView');

	var calendar = new FullCalendar.Calendar(calendarEl, {
		plugins: [ 'list' ],

		header: {
			left: 'prev,next today',
			center: 'title',
			right: 'listDay,listWeek,dayGridMonth'
		},

		// customize the button names,
		// otherwise they'd all just say "list"
		views: {
			listDay: { buttonText: 'list day' },
			listWeek: { buttonText: 'list week' }
		},

		defaultView: 'listWeek',
		defaultDate: '2019-12-12',
		navLinks: true, // can click day/week names to navigate views
		editable: true,
		eventLimit: true, // allow "more" link when too many events
		events: [
			{
				title: 'All Day Event',
				start: '2019-12-01'
			},
			{
				title: 'Long Event',
				start: '2019-12-07',
				end: '2019-12-10'
			},
			{
				groupId: 999,
				title: 'Repeating Event',
				start: '2019-12-09T16:00:00'
			},
			{
				groupId: 999,
				title: 'Repeating Event',
				start: '2019-12-16T16:00:00'
			},
			{
				title: 'Conference',
				start: '2019-12-11',
				end: '2019-12-13'
			},
			{
				title: 'Meeting',
				start: '2019-12-12T10:30:00',
				end: '2019-12-12T12:30:00'
			},
			{
				title: 'Lunch',
				start: '2019-12-12T12:00:00'
			},
			{
				title: 'Meeting',
				start: '2019-12-12T14:30:00'
			},
			{
				title: 'Happy Hour',
				start: '2019-12-12T17:30:00'
			},
			{
				title: 'Dinner',
				start: '2019-12-12T20:00:00'
			},
			{
				title: 'Birthday Party',
				start: '2019-12-13T07:00:00'
			},
			{
				title: 'Click for Dashboard',
				url: 'index.html',
				start: '2019-12-28'
			}
		]
	});

	calendar.render();
});