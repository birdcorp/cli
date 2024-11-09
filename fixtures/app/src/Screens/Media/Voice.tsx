import React, { useEffect, useState } from 'react'
import { HighlightedCode } from '../../Components'

import { Mic } from 'lucide-react'

export default function Voice() {
  const [text, setText] = useState('')

  const code = `window.miniProgram?.call(
      'openVoiceInput',
      {
        "suggestions": [
         
        ]
      },
      {
        success(response) {
          if (response.ok) {

          }
        }
      },
    )`

  function trigger() {
    window.miniProgram?.call(
      'openVoiceInput',
      {
        commands: [
          {
            comand: 'search',
            description: 'Search product comand',
          },
        ],
      },
      {
        success(response: any) {
          if (response.ok) {
            setText(response.result)

            /*
            switch (response.result.comand) {
              case 'search':
                setText('Search product comand')
                break
              default:
                setText('Unknown comand')
                break
            }
                */
          }
        },
      },
    )
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Open Voice Input',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Open Voice Input</h3>

      <p className="text-gray-300 mb-8">Opens voice to text input</p>

      {text && <p className="text-gray-100 italic text-4xl mb-10">"{text}"</p>}

      <button
        onClick={trigger}
        className="text-white font-bold py-2 px-4 rounded"
      >
        <Mic />
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}
