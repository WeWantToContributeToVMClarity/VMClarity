import React from 'react';
import { useLocation } from 'react-router-dom';
import TabbedPage from 'components/TabbedPage';
import FindingsDetailsPage from '../FindingsDetailsPage';
import TabRootkitDetails from './TabRootkitDetails';

const ROOTKIT_DETAILS_PATHS = {
    PACKAGE_DETAILS: "",
}

const DetailsContent = ({data}) => {
    const {pathname} = useLocation();
    
    const {id} = data;
    
    return (
        <TabbedPage
            basePath={`${pathname.substring(0, pathname.indexOf(id))}${id}`}
            items={[
                {
                    id: "general",
                    title: "Rootkit details",
                    isIndex: true,
                    component: () => <TabRootkitDetails data={data} />
                }
            ]}
            withInnerPadding={false}
        />
    )
}

const RootkitDetails = () => (
    <FindingsDetailsPage
        backTitle="Rootkits"
        getTitleData={({findingInfo}) => ({title: findingInfo.rootkitName})}
        detailsContent={DetailsContent}
    />
)

export default RootkitDetails;
