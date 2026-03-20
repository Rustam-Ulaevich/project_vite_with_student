import audio1 from '../audio/viktor-cojj-zvezda-po-imeni-solnce.mp3'
import audio2 from '../audio/abba-happy-new-year.mp3'
import audio3 from '../audio/ANNA_ASTI_-_Po_baram.mp3'
import { useEffect, useState } from 'react'

export const Content = () => {

	const[traсks, setTracks] = useState([
		{id: 1, title: 'Viktor Cojj', genre: 'Rock', url: audio1},
		{id: 2, title: 'Abba', genre: 'Hit80', url: audio2},
		{id: 3, title: 'Anna Asti', genre: 'Pop', url: audio3},
	])
	

	useEffect( () => {
        (async () => {
            try {
                const data = await fetch('https://62c3ffff7d83a75e39ecd122.mockapi.io/orders') // Необходимо создать собственный подобный сервис Mockapi.io
								.then(res => res.json())
								.then(json => console.log(json))
                
            } catch (error) {
                alert('Error when requesting orders')
                console.log(error)
            }
        })();
    }, [])

	// fetch('https://api.ittybit.com/tasks', {
  // 	method: 'POST',
  // 	headers: {
  //   	'Authorization': 'Bearer <ITTYBIT_API_KEY>',
  //   	'Content-Type': 'application/json',
  // 	},
  // 	body: JSON.stringify({
  //   	url: 'https://webapp.ittybit.net/homepage-dance.mp4',
  //   	kind: 'audio',
  //   	options: { format: 'mp3' },
  // 	})
	// }).then(res => res.json())
	// .then(json => console.log(json))
	
	

	if(traсks.length === 0){
		return(
			<div  className="content">NO TRACK!</div>
		)
	}


    return (
      <>
				<div className="content">
					<div>
						<button>Rock</button>
						<button>Pop</button>
						<button>Hit80</button>
					</div>

					<ul>
					{traсks.map( traсk => (
							<li className="track" key={traсk.id}>
								<div className="track_discription">
									<p>{traсk.title}</p>
									<p>{traсk.genre}</p>
								</div>

								<audio controls src={traсk.url}></audio>
							</li>
						))}

					</ul>
				</div>
				
			</>
    )
}