import NotLoggedNavbar from "@/components/Navbar/NotLoggedIn";

export default function Team() {
    return (
        <div className="justify-center items-center min-h-screen bg-gradient-to-r from-gray-800 to-gray-900">
            <NotLoggedNavbar />
            Team Page
        </div>
    );
}