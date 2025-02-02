import React, { Fragment, useEffect } from "react";

import {
  EuiPageTemplate,
  EuiSpacer,
} from "@elastic/eui";
import { FormContextProvider, replaceBreadcrumbs } from "@gojek/mlp-ui";
import { useNavigate, useParams } from "react-router-dom";

import { PageTitle } from "components/page/PageTitle";
import { SegmentContextProvider } from "providers/segment/context";
import { SegmenterContextProvider } from "providers/segmenter/context";
import { EditSegmentForm } from "segments/components/form/EditSegmentForm";
import { CustomSegment } from "services/segment/CustomSegment";

const EditSegmentView = ({ segmentSpec }) => {
  const { projectId } = useParams();
  const navigate = useNavigate();

  useEffect(() => {
    replaceBreadcrumbs([
      { text: "Experiments", href: "../.." },
      { text: "Segments", href: ".." },
      { text: segmentSpec.name, href: "." },
      { text: "Configuration" },
    ]);
  });

  return (
    <Fragment>
      <EuiPageTemplate.Header
        bottomBorder={false}
        pageTitle={<PageTitle title="Edit Segment" />}
      />
      <EuiSpacer size="l" />
      <EuiPageTemplate.Section color={"transparent"}>
        <FormContextProvider data={CustomSegment.fromJson(segmentSpec)}>
          <SegmenterContextProvider projectId={projectId} status="active">
            <SegmentContextProvider projectId={projectId}>
              <EditSegmentForm
                projectId={projectId}
                onCancel={() => window.history.back()}
                onSuccess={() => {
                  navigate("../", { state: { refresh: true } });
                }}
              />
            </SegmentContextProvider>
          </SegmenterContextProvider>
        </FormContextProvider>
      </EuiPageTemplate.Section>
    </Fragment>
  );
};

export default EditSegmentView;
