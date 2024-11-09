import React, { useState } from 'react'

// import { ShippingAddress } from 'birdcash-miniprogram-sdk-alpha'

export interface ShippingAddress {
  id: string
  firstName?: string
  lastName?: string
  addressLine1?: string
  addressLine2?: string
  city?: string
  state?: string
  zipCode?: string
  country?: string
}

const ShippingAddressForm: React.FC = () => {
  const [address, setAddress] = useState<ShippingAddress>({
    id: new Date().toISOString(),
    firstName: '',
    lastName: '',
    addressLine1: '',
    addressLine2: '',
    city: '',
    state: '',
    zipCode: '',
    country: '',
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setAddress({ ...address, [e.target.name]: e.target.value })
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    console.log('Shipping Address:', address)
    // Add logic for submitting the form data, like making an API call
  }

  return (
    <form className="" onSubmit={handleSubmit}>
      <div className="flex space-x-4 mb-4">
        <div className="w-1/2">
          <label className="block text-sm font-medium mb-1 text-gray-300">
            First Name
          </label>
          <input
            type="text"
            name="firstName"
            value={address.firstName}
            onChange={handleChange}
            className="w-full border bg-neutral-800 border-gray-700 p-2 rounded text-white"
          />
        </div>
        <div className="w-1/2">
          <label className="block text-sm font-medium mb-1 text-gray-300">
            Last Name
          </label>
          <input
            type="text"
            name="lastName"
            value={address.lastName}
            onChange={handleChange}
            className="w-full border bg-neutral-800 border-gray-700 p-2 rounded text-white"
          />
        </div>
      </div>

      <div className="flex space-x-4 mb-4">
        <div className="w-1/2">
          <label className="block text-sm font-medium mb-1 text-gray-300">
            Address Line 1
          </label>
          <input
            type="text"
            name="addressLine1"
            value={address.addressLine1}
            onChange={handleChange}
            className="w-full border bg-neutral-800 border-gray-700 p-2 rounded text-white"
          />
        </div>
        <div className="w-1/2">
          <label className="block text-sm font-medium mb-1 text-gray-300">
            Address Line 2
          </label>
          <input
            type="text"
            name="addressLine2"
            value={address.addressLine2}
            onChange={handleChange}
            className="w-full border bg-neutral-800 border-gray-700 p-2 rounded text-white"
          />
        </div>
      </div>

      <div className="flex space-x-4 mb-4">
        <div className="w-1/2">
          <label className="block text-sm font-medium mb-1 text-gray-300">
            City
          </label>
          <input
            type="text"
            name="city"
            value={address.city}
            onChange={handleChange}
            className="w-full border bg-neutral-800 border-gray-700 p-2 rounded text-white"
          />
        </div>
        <div className="w-1/2">
          <label className="block text-sm font-medium mb-1 text-gray-300">
            State
          </label>
          <input
            type="text"
            name="state"
            value={address.state}
            onChange={handleChange}
            className="w-full border bg-neutral-800 border-gray-700 p-2 rounded text-white"
          />
        </div>
      </div>

      <div className="flex space-x-4 mb-4">
        <div className="w-1/2">
          <label className="block text-sm font-medium mb-1 text-gray-300">
            ZIP Code
          </label>
          <input
            type="text"
            name="zipCode"
            value={address.zipCode}
            onChange={handleChange}
            className="w-full border bg-neutral-800 border-gray-700 p-2 rounded text-white"
          />
        </div>
        <div className="w-1/2">
          <label className="block text-sm font-medium mb-1 text-gray-300">
            Country
          </label>
          <input
            type="text"
            name="country"
            value={address.country}
            onChange={handleChange}
            className="w-full border bg-neutral-800 border-gray-700 p-2 rounded text-white"
          />
        </div>
      </div>
    </form>
  )
}

export default ShippingAddressForm
