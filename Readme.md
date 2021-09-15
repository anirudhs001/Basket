<h1>Basket </h1>  
A basic website written in golang(mostly) and JS to ease offline shopping.

To see in action, run the executable as: ./Basket and goto localhost:8080

<br>
<h2>What's working</h2>
<h3>User session</h3>
user data is stored in a cookie to mantain persistance across different pages
<h3>storage</h3>
user's items-lists and seller data is stored across multiple postgres tables
<h2>What needs work</h2>

<ul>
    <li>Improve the UI</li>
    <li>seller registration page</li>
        the backend's done; but it's still a simple page with just 4 input boxes
    <li>customer can't view available items in shop</li>
        customers can't yet view the items which were available at the shop

</ul>
<h2>How to use</h2>
<ul>
    <li> <h3> Buyers </h3> </li>
    <h4>Basket</h4>
    <p>
    multiple users(family members, people in the same building) can add items to the
    same basket, which can be sent to the local shops. This would allow customers to shop easily without waiting in line and mantain distancing.
    </p>
    <h4>Shops around you</h4>
    <p>All nearby shops who registered on the website are visible here</p>
    <li> <h3> Sellers </h3> </li>
    <h4> register </h4>
    Sellers can register their shop address and timings.
    <h4>
    View Items
    </h4>
    Seller can view all the requested items and pack them early on.
    </ul>
<h4>note</h4>

<i>contents in the public folder are accessible by anyone<i>
