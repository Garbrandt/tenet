// Week Numbers *******************
document.addEventListener('DOMContentLoaded', function() {
	var calendarEl = document.getElementById('calendarWeekNumbers');

	var calendar = new FullCalendar.Calendar(calendarEl, {
		plugins: [ 'interaction', 'dayGrid', 'timeGrid', 'list' ],
		header: {
			left: 'prev,next today',
			center: 'title',
			right: 'dayGridMonth,timeGridWeek,timeGridDay,listWeek'
		},
		navLinks: true, // can click day/week names to navigate views

		weekNumbers: true,
		weekNumbersWithinDays: true,
		weekNumberCalculation: 'ISO',

		editable: true,
		eventLimit: true, // allow "more" link when too many events
		events: [
			{
				title: 'All Day Event',
				start: '2019-08-01'
			},
			{
				title: 'Long Event',
				start: '2019-08-07',
				end: '2019-08-10'
			},
			{
				groupId: 999,
				title: 'Repeating Event',
				start: '2019-08-09T16:00:00'
			},
			{
				groupId: 999,
				title: 'Repeating Event',
				start: '2019-08-16T16:00:00'
			},
			{
				title: 'Conference',
				start: '2019-08-11',
				end: '2019-08-13'
			},
			{
				title: 'Meeting',
				start: '2019-08-12T10:30:00',
				end: '2019-08-12T12:30:00'
			},
			{
				title: 'Lunch',
				start: '2019-08-12T12:00:00'
			},
			{
				title: 'Meeting',
				start: '2019-08-12T14:30:00'
			},
			{
				title: 'Happy Hour',
				start: '2019-08-12T17:30:00'
			},
			{
				title: 'Dinner',
				start: '2019-08-12T20:00:00'
			},
			{
				title: 'Birthday Party',
				start: '2019-08-13T07:00:00'
			},
			{
				title: 'Click for Dashboard',
				url: 'index.com',
				start: '2019-08-28'
			}
		]
	});

	calendar.render();
});