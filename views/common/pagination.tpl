<style>
    /* 自定义分页样式 */
    .pagination-wrapper {
        margin: 20px 0;
        text-align: center;
    }

    .pagination {
        display: inline-block;
        padding: 0;
        margin: 0;
    }

    .pagination li {
        display: inline;
        margin: 0 4px;
    }

    .pagination li a, .pagination li span {
        color: #337ab7;
        padding: 8px 16px;
        text-decoration: none;
        border: 1px solid #ddd;
        border-radius: 4px;
    }

    .pagination li.active span {
        background-color: #337ab7;
        color: white;
        border: 1px solid #337ab7;
    }

    .pagination li.disabled span {
        color: #ddd;
    }

    .pagination li a:hover {
        background-color: #eee;
    }
</style>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">

<div class="pagination-wrapper text-center">
    {{ if .Pagination }}
    <nav aria-label="Page navigation">
        {{ .Pagination.PageLinks }}
    </nav>
    {{ end }}
</div>
