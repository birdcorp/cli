import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function NavigationBar() {
  const code = `window.miniProgram?.call("setNavigationBar", {
  title: 'New Title',
  backgroundColor: '#009cda',
  color: 'light',
})`

  function trigger() {
    window.miniProgram?.call('setNavigationBar', {
      title: getRandomTitle(),
      backgroundColor: getRandomHex(),
      color: 'light',
    })
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Set Navigation Bar',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 dark:bg-black dark:text-gray-100">
      <h3 className="text-white font-bold mb-3">Update NavigationBar</h3>

      <p className="text-gray-400 mb-8 dark:text-gray-300">
        Updates the navigation bar title, text color, and background color
      </p>

      <button
        onClick={trigger}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
      >
        Update
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-400 dark:text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />
    </div>
  )
}

function getRandomHex(): string {
  // Generates a random hex color
  const randomColor = Math.floor(Math.random() * 16777215)
    .toString(16)
    .padStart(6, '0')
  const hexColor = `#${randomColor}`

  // Convert hex to RGB
  const rgb = parseInt(randomColor, 16)
  const r = (rgb >> 16) & 0xff
  const g = (rgb >> 8) & 0xff
  const b = rgb & 0xff

  // Calculate luminance
  const luminance = 0.299 * r + 0.587 * g + 0.114 * b

  // Ensure the color is dark enough to have good contrast with white text
  // A luminance below 160 is considered a dark enough background
  if (luminance > 160) {
    return getRandomHex() // Recursively generate a new color if it's too light
  }

  return hexColor
}

function getRandomTitle() {
  const titles = [
    'Home',
    'Search',
    'Profile',
    'Settings',
    'Notifications',
    'Messages',
    'Favorites',
    'Explore',
    'Help',
    'About',
    'Contact',
    'Feedback',
    'Terms of Service',
    'Privacy Policy',
    'Dashboard',
  ]

  // Get a random index based on the length of the array
  const randomIndex = Math.floor(Math.random() * titles.length)

  return titles[randomIndex]
}
