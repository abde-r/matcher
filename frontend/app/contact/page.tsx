import NotLoggedNavbar from "@/components/Navbar/NotLoggedIn";
import ContactComponent from "@/components/Contact/contact";

export default function Contact() {
    return (
        <div className="justify-center items-center min-h-screen bg-gradient-to-r from-gray-800 to-gray-900">
            <NotLoggedNavbar />
            <ContactComponent />
        </div>
    );
}
