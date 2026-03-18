
import logo from '../assets/logo.png'

export const Navbar = () => {
    return (
			<>
				<nav className='navbar'>

        <div className="logo">
          <img src={logo}  width='35px' height='35px' alt="Logo"/>
          <h1>Musicfun</h1>
        </div>
        
        <ul>
          <li><a href="http://">Main</a></li>
          <li><a href="http://">Favorit</a></li>
          <li><a href="http://">Contacts</a></li>
          <li><a href="http://">Orders</a></li>
        </ul>      
      </nav>
			</>
		)
}