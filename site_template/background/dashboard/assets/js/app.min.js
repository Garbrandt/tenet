!function (e) {
    "use strict";

    function t() {
    }

    t.prototype.initTooltipPlugin = function () {
        e.fn.tooltip && e('[data-toggle="tooltip"]').tooltip()
    }, t.prototype.initPopoverPlugin = function () {
        e.fn.popover && e('[data-toggle="popover"]').popover()
    }, t.prototype.initSlimScrollPlugin = function () {
        e.fn.slimScroll && e(".slimscroll").slimScroll({
            height: "auto",
            position: "right",
            size: "4px",
            touchScrollStep: 20,
            color: "#9ea5ab"
        })
    }, t.prototype.initFormValidation = function () {
        e(".needs-validation").on("submit", function (t) {
            return e(this).addClass("was-validated"), !1 !== e(this)[0].checkValidity() || (t.preventDefault(), t.stopPropagation(), !1)
        })
    }, t.prototype.init = function () {
        this.initTooltipPlugin(), this.initPopoverPlugin(), this.initSlimScrollPlugin(), this.initFormValidation()
    }, e.Components = new t, e.Components.Constructor = t
}(window.jQuery), function (n) {
    "use strict";

    function t() {
        this.$body = n("body"), this.$window = n(window)
    }

    t.prototype._resetSidebarScroll = function () {
        n(".slimscroll-menu").slimscroll({
            height: "auto",
            position: "right",
            size: "4px",
            color: "#9ea5ab",
            wheelStep: 5,
            touchScrollStep: 20
        })
    }, t.prototype.initMenu = function () {
        var e = this;
        n(".button-menu-mobile").on("click", function (t) {
        });
        var t = e.$body.data("layout");
        if (n("#menu-bar").length) if ("topnav" !== t) new MetisMenu("#menu-bar"), e._resetSidebarScroll(), n("#menu-bar a").each(function () {
            var t = window.location.href.split(/[?#]/)[0];
            this.href == t && (n(this).addClass("active"), n(this).parent().addClass("mm-active"), n(this).parent().parent().addClass("mm-show"), n(this).parent().parent().prev().addClass("active"), n(this).parent().parent().parent().addClass("mm-active"), n(this).parent().parent().parent().parent().addClass("mm-show"), n(this).parent().parent().parent().parent().parent().addClass("mm-active"))
        }); else {
            var i = new MetisMenu("#menu-bar").on("shown.metisMenu", function (n) {
                window.addEventListener("click", function t(e) {
                    n.target.contains(e.target) || (i.hide(n.detail.shownElement), window.removeEventListener("click", t))
                })
            });
            n("#menu-bar a").each(function () {
                var t = window.location.href.split(/[?#]/)[0];
                this.href == t && (n(this).addClass("active"), n(this).parent().addClass("active"), n(this).parent().parent().prev().addClass("active"), n(this).parent().parent().parent().addClass("active"), n(this).parent().parent().parent().parent().parent().addClass("active"))
            })
        }
        n(".right-bar-toggle").on("click", function (t) {
            n("body").toggleClass("right-bar-enabled")
        }), n(document).on("click", "body", function (t) {
        }), n(window).on("load", function () {
            n("#status").fadeOut(), n("#preloader").delay(350).fadeOut("slow")
        })
    }, t.prototype.initLayout = function () {
    }, t.prototype.init = function () {
        var e = this;
        this.initLayout(), this.initMenu(), n.Components.init(), e.$window.on("resize", function (t) {
            t.preventDefault(), e.initLayout(), e._resetSidebarScroll()
        }), feather.replace()
    }, n.App = new t, n.App.Constructor = t
}(window.jQuery), function () {
    "use strict";
    window.jQuery.App.init()
}();
//# sourceMappingURL=app.min.js.map
