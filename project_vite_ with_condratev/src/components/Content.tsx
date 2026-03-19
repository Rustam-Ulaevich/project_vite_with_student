import audio1 from '../audio/viktor-cojj-zvezda-po-imeni-solnce.mp3'
import audio2 from '../audio/abba-happy-new-year.mp3'
import audio3 from '../audio/ANNA_ASTI_-_Po_baram.mp3'

export const Content = () => {
	
	const traсks = [
		{id: 1, title: 'Viktor Cojj', genre: 'Rock', url: audio1},
		{id: 2, title: 'Abba', genre: 'Hit80', url: audio2},
		{id: 3, title: 'Anna Asti', genre: 'Pop', url: audio3},
	]

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