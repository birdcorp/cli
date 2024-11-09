import React from 'react'
import { HighlightedCode } from '../../Components'

export default function Models() {
  const code = `
  <a
  className="block my-4"
  href="https://example.com/models/pancakes.usdz"
  rel="ar"
>
  Pancake
  <img
    className="w-60 h-60"
    src="https://example.com/models/pancakes_thumbnail.png"
    alt="View in AR"
  />
</a>
  `

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg mx-auto px-3 pt-5">
      <h3 className="text-white font-bold mb-3">AR Models</h3>

      <p className="text-gray-400 mb-8">View 3D models in augmented reality.</p>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4 p-4">
        <a
          className="block relative bg-neutral-800 rounded-lg shadow-lg hover:shadow-xl transition-shadow duration-200 my-4 overflow-hidden"
          href="https://dlkosrb2bmrzf.cloudfront.net/models/pancakes.usdz"
          rel="ar"
        >
          <img
            className="w-full h-60 object-cover"
            src="https://dlkosrb2bmrzf.cloudfront.net/models/pancakes_thumbnail.png"
            alt="View in AR"
          />
          <div className="absolute bottom-0 w-full bg-black bg-opacity-50 text-white text-lg font-semibold p-2 text-center">
            Pancake
          </div>
        </a>

        <a
          className="block relative bg-neutral-800 rounded-lg shadow-lg hover:shadow-xl transition-shadow duration-200 my-4 overflow-hidden"
          href="https://dlkosrb2bmrzf.cloudfront.net/models/sneaker_airforce.usdz"
          rel="ar"
        >
          <img
            className="w-full h-60 object-cover"
            src="https://dlkosrb2bmrzf.cloudfront.net/models/sneaker_airforce_thumbnail.png"
            alt="View in AR"
          />
          <div className="absolute bottom-0 w-full bg-black bg-opacity-50 text-white text-lg font-semibold p-2 text-center">
            Sneaker Airforce
          </div>
        </a>

        <a
          className="block relative bg-neutral-800 rounded-lg shadow-lg hover:shadow-xl transition-shadow duration-200 my-4 overflow-hidden"
          href="https://dlkosrb2bmrzf.cloudfront.net/models/sneaker_pegasustrail.usdz"
          rel="ar"
        >
          <img
            className="w-full h-60 object-cover"
            src="https://dlkosrb2bmrzf.cloudfront.net/models/sneaker_pegasustrail_thumbnail.png"
            alt="View in AR"
          />
          <div className="absolute bottom-0 w-full bg-black bg-opacity-50 text-white text-lg font-semibold p-2 text-center">
            Sneaker Pegasustrail
          </div>
        </a>
      </div>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-400">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
