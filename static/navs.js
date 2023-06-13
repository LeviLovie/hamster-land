document.body.innerHTML += `
<nav class=\"horiz-nav\">
    <!-- <div class=\"horiz-nav-element\" style=\"margin: 10px 30px 0 10px;\"><h5>Home</h5></div> -->
    <div class=\"horiz-nav-element\" style=\"margin: 10px 30px 0 10px;\"><h5>17 Cº, Rainy day</h5></div>
    <div class=\"horiz-nav-element\" style=\"margin: 10px 30px 0 10px;\"><h5>Creating the site!</h5></div>
</nav>
<nav class=\"vert-nav\">
    <div class=\"vert-nav-buttons\">
        <div class=\"vert-nav-button\" onclick=\"window.location.href='/'\"><h5 class=\"vert-nav-button-text\">Home</h5></div>
        <div class=\"vert-nav-button\" onclick=\"window.location.href='/map'\"><h5 class=\"vert-nav-button-text\">Map</h5></div>
        <div class=\"vert-nav-button\" onclick=\"window.location.href='/bank'\"><h5 class=\"vert-nav-button-text\">Bank</h5></div>
        <div class=\"vert-nav-button\" onclick=\"window.location.href='/news'\"><h5 class=\"vert-nav-button-text\">News</h5></div>
        <div class=\"vert-nav-button\" onclick=\"window.location.href='/about'\"><h5 class=\"vert-nav-button-text\">About</h5></div>
    </div>
</nav>
`