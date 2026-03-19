
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
          <li><a href="http://google.com">Main</a></li>
          <li><a href="http://xn----7sbnjgaaco2agjzw7hxb.xn--p1ai/vacancies/afrikakorpus?utm_source=yandex&utm_medium=cpc&utm_campaign=707719711&mango=|c:707719711|g:5724129990|b:1903178616426774918|k:205724129990|st:context|a:no|s:rus.hitmotop.com|t:none|p:0|r:205724129990|reg:10761|net:{yad}&yclid=3460780122883751935">Favorit</a></li>
          <li><a href="http://">Contacts</a></li>
          <li><a href="http://">Orders</a></li>
        </ul>      
      </nav>
			</>
		)
}