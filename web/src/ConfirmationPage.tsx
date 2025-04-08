import { useNavigate, useParams } from "react-router-dom"
import { API_URL } from "./App"

export const ConfirmationPage = () => {
    const {token = ""} = useParams()
    const redirect = useNavigate()
    
    const handleConfirm = async () => {
        const response = await fetch(`${API_URL}/users/activate/${token}`, {
            method: "PUT"
        })
        console.log(response)

        if (response.ok) {
            // redirect to home
            redirect("/")
        } else {
            // handle error
            alert(`Failed to confirm token, ${response.statusText}`)
        }
    }

    return (
        <div>
            <h1>Confirmation</h1>
            <button onClick={handleConfirm}>Click here to confirm</button>
        </div>
    )
}