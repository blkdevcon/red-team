function removeEmpty(e, page_id) {
    if (e.text().trim().length == 0) {
        console.debug("Removing empty element from " + page_id + ": ");
        console.debug(e);

        p = e.parent();
        e.remove();

        if (p.is("table")) {
            let c = p.clone(), h = c.children("thead");
            h.remove();
            if (c.text().trim().length == 0) {
                console.debug("Removing empty table from " + page_id + ": ");
                console.debug(p);
                p.remove();
            }
        }
    }

}

let last_node = false, last_thead = false, last_tr = false;
class ElementCleaner extends Paged.Handler {

    constructor(chunker, polisher, caller) {
        super(chunker, polisher, caller);
    }

    afterParsed(parsed) {
        $("body#report-body > div.pagedjs_pages").css("justify-content", "center");
    }

    afterPageLayout(pageFragment, page) {
        const cleanup = [ 
            ".generated-table table", 
            ".finding-content pre",
            ".finding-content",
            "tbody"
        ];

        cleanup.forEach( e => $(page.element).find(e).each( function() { removeEmpty($(this), page.id); } ) );
    }

    afterRendered(pages) {
        let t1 = performance.now();
        console.log("Rendering took " + Number.parseFloat((t1 - t0)/1000).toPrecision(3) + " seconds.");

        $("<input type='hidden' value='1' id='finished_loading' />").appendTo("html");
        console.log($("#finished_loading"));
    }

    renderNode(node, sourceNode, layout) {
        // Handle numbered lists that are split by code blocks
        if ($(node).is("pre") && $(node).prev().is("ol")) {
            let p = $(node).prev();
            let l = p.children("li:last-child");
            $(node).appendTo(l);
        } else if ($(node).is("li") && $(node).parent().is("ol") && $(node).parent().prev().is("ol")) {
            let p_orig = $(node).parent();
            let p = $(node).parent().prev();
            $(node).appendTo(p);
            if ($(sourceNode).is(":last-child")) {
                p_orig.remove();
            }
        }

        if (last_node) {
            // Special handling for split tables: clone sibling <td> elements
            if (last_tr) {
                if ($(node).parent().is("td")) {
                    let split_from = $(node).parent().attr("data-split-from");
                    let last_td = $(last_tr).children("td[data-ref=" + split_from + "]");
                    let new_td = $(node).closest("td");
                    let new_tbody = new_td.closest("tbody");

                    // Add split content back to prior <td>, then append the
                    // prior <td> to the new table
                    last_td.html(last_td.html() + new_td.html());
                    new_tbody.html(last_tr);
                    last_tr = false;

                    $(new_tbody).closest("table").prepend(last_thead.clone());
                    last_thead = false;
                }
            }

            let comp = last_node;
            if ($(node).is("tbody")) {
                comp = $(node).parent();
            }

            if ($(last_node).attr("data-ref") == $(comp).attr("data-ref")) {
                if (last_thead) {
                    console.debug("Prepending last thead to split table");
                    $(node).closest("table").prepend(last_thead.clone());
                    last_thead = false;
                }

                console.debug($(node).closest("table"));
                if ($(last_node).is("td") || $(last_node).parent().is("td")) {
                    let tr = $(last_node).closest("tr"), tbody = tr.parent();
                    if (tr.text().trim().length == 0) {
                        console.debug("Removing empty row from split table");
                        tr.remove();
                        tbody.css("transform", "translateZ(0)"); // Force CSS redraw
                    }
                }
                last_node = false;
            }
        }
    }

    onBreakToken(breakToken, overflow, rendered) {
        if (overflow) {
            let sc = $(overflow.startContainer)[0];
            if ($(sc).is("td") || $(sc).parent().is("td") || $(sc).is("tbody")) {
                let table = $(sc).closest("table");
                if (table.parent().is(".generated-table")) {
                    console.debug("Page break in .generated-table; storing last thead element");
                    last_thead = $(table).children("thead");
                    if ($(sc).is("tbody")) {
                        last_node = table;
                    } else {
                        last_node = sc;
                        last_tr = sc.closest("tr");
                    }
                }
            }
        }
    }
}

Paged.registerHandlers(ElementCleaner);

let t0 = performance.now();
$().ready( function() { 
    window.PagedPolyfill.preview();
});
