// 一个新消息sider
var newSiderItem = `<li class="list-group-item">
    <figure class="avatar">
    <span class="avatar-title bg-warning bg-success rounded-circle">
        <i class="fa fa-users"></i>
    </span>
    </figure>
    <div class="users-list-body">
        <div>
            <h5>😍 My Family 😍</h5>
            <p><strong>Maher Ruslandi: </strong>Hello!!!</p>
        </div>
        <div class="users-list-action">
            <small class="text-muted">Yesterday</small>
            <div class="action-toggle">
                <div class="dropdown">
                    <a data-toggle="dropdown" href="#">
                        <i data-feather="more-horizontal"></i>
                    </a>
                    <div class="dropdown-menu dropdown-menu-right">
                        <a href="#" class="dropdown-item">Open</a>
                        <a href="#" data-navigation-target="contact-information"
                            class="dropdown-item">Profile</a>
                        <a href="#" class="dropdown-item">Add to archive</a>
                        <div class="dropdown-divider"></div>
                        <a href="#" class="dropdown-item text-danger">Delete</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</li > `

function newSideChatItem() {
    var group = document.getElementById("sidebar_item_group")
    group.insertAdjacentHTML("beforeEnd", newSiderItem)
}