import audio1 from '../audio/viktor-cojj-zvezda-po-imeni-solnce.mp3'

export const Content = () => {

	const traсks = [
		{id: 1, title: 'Pop', genre: 'Pop', url: {audio1}},
		{id: 2, title: 'Rock', genre: 'Rock', url: ''},
		{id: 3, title: 'Rock', genre: 'Rock', url: ''},
		{id: 4, title: 'Rock', genre: 'Rock', url: ''}
	]

    return (
      <>
			<div className="content">
				<ul>
					{traсks.map( traсk => (
						<li className="track" key={traсk.id}>
							<div className="track_discription">
								<p>{traсk.title}</p>
								<p>{traсk.genre}</p>
							</div>
							
							<audio controls src='{traсk.url}'></audio>
						</li>
					))}

				</ul>
			</div>
				
			</>
    )
}