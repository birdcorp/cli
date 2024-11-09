import React, { useEffect, useState } from 'react'
import { HighlightedCode } from '../../Components'

export default function ShareLink() {
  function trigger() {
    window.miniProgram?.call(
      'share',
      {
        path: '/share/link',
        title: 'Product Title',
        description: 'some descripton',
        image:
          'https://flowbite.s3.amazonaws.com/docs/gallery/square/image-4.jpg',
      },
      {
        success(response) {
          if (response.ok) {
          }
        },
        failure(err) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Share Link',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Share Link</h3>

      <p className="text-gray-300 mb-8">Shares a link</p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Share
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call(
      'share',
      {
        path: '/share/link',
        title: 'Product Title',
        description: 'some descripton',
        image:
          'https://flowbite.s3.amazonaws.com/docs/gallery/square/image-4.jpg',
      },
      {
        success(response) {
          if (response.ok) {
          }
        },
        failure(err) {
          console.error(err)
        },
        completed() {},
      },
    )`}
      />
    </div>
  )
}
