import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function Video() {
  const code = `window.miniProgram?.call("openVideo", {
    title: "My Video",
    url: "https://www.example.com/video/city.mp4"
})`

  function trigger() {
    window.miniProgram?.call('openVideo', {
      title: 'My Video',
      url: 'https://dlkosrb2bmrzf.cloudfront.net/video/city.mp4',
    })
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Open Video',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Open Video</h3>

      <p className="text-gray-300 mb-8">
        Opens a full screen video screen sheet
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Call
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
