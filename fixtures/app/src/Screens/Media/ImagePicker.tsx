import React, { useEffect, useState } from 'react'
import { HighlightedCode } from '../../Components'

export default function ImagePicker() {
  const code = `window.miniProgram?.call('selectImage', {
    title: 'Select Images',
})`

  const [response, setResponse] = useState({
    filePaths: [],
  })

  function trigger() {
    window.miniProgram?.call(
      'selectImages',
      {},
      {
        success(response: any) {
          setResponse(response)
          if (response.ok) {
            window.miniProgram?.call(
              'uploadFiles',
              {
                url: 'https://httpbin.org/post',
                filePaths: response.filePaths,
                names: (response.filePaths as string[]).map((filePath) => {
                  return filePath.split('/').pop()
                }),
                formData: {
                  key: 'value',
                },
              },
              {
                success(response: any) {
                  // setResponse(response)

                  if (response.ok) {
                  }
                },
              },
            )
          }
        },
        failure(err: any) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Image Picker',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Image Picker</h3>

      <pre>
        <code>{JSON.stringify(response, null, 4)}</code>
      </pre>

      <p className="text-gray-300 mb-8">Selects images</p>

      <h3 className="text-white font-bold mb-3">Example</h3>

      <p className="text-gray-300 mb-8">
        Selects images and calls uploadFiles to upload them to a server
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Select
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call(
      'selectImages',
      {},
      {
        success(response) {
          if (response.ok) {

            window.miniProgram?.call(
              'uploadFiles',
              {
                url: 'https://your-server.com/upload',
                filePaths: response.filePaths,
                names: response.filePaths.map((filePath) => 
                  filePath.split('/').pop()),
                formData: {
                  key: 'value',
                },
              },
              {
                success(response) {
                  if (response.ok) {

                  }
                },
              },
            )
              
          }
        },
      },
    )`}
      />
    </div>
  )
}
